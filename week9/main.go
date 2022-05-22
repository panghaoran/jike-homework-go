package main

import (
	"fmt"
)

func main(){
    
}

/*
* goim 协议结构
* 4bytes PacketLen
* 2bytes HeaderLen
* 2bytes Version
* 4bytes Operation
* 4bytes Sequence
* PacketLen-HeaderLen Body
*/
func (bigEndian bigEndian) WriteTo(body string) []byte {
	packLen := _rawHeaderSize + len(body)
	ret := make([]byte, packLen)

	bigEndian.PutInt32(ret[_packOffset:], int32(packLen))
	bigEndian.PutInt16(ret[_headerOffset:], int16(_rawHeaderSize))

	version := 1
	bigEndian.PutInt16(ret[_verOffset:], int16(version))
	operation := 1
	bigEndian.PutInt32(ret[_opOffset:], int32(operation))
	sequence := 1
	bigEndian.PutInt32(ret[_seqOffset:], int32(sequence))
	
	byteBody := []byte(body)
	copy(ret[_rawHeaderSize:], byteBody)
	return ret
}

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

/*
* goim 协议结构
* 4bytes PacketLen
* 2bytes HeaderLen
* 2bytes Version
* 4bytes Operation
* 4bytes Sequence
* PacketLen-HeaderLen Body
*/
func (bigEndian bigEndian) Read(data []byte) {
	packLen := bigEndian.Int32(data[_packOffset:_headerOffset])
	fmt.Printf("packetLen:%v\n", packLen)

	headerLen := bigEndian.Int16(data[_headerOffset:_verOffset])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := int32(bigEndian.Int16(data[_verOffset:_opOffset]))
	fmt.Printf("version:%v\n", version)

	operation := bigEndian.Int32(data[_opOffset:_seqOffset])
	fmt.Printf("operation:%v\n", operation)

	sequence := bigEndian.Int32(data[_seqOffset:])
	fmt.Printf("sequence:%v\n", sequence)

	// bodyLen = int(packLen - int32(headerLen)); bodyLen > 0 {
	// p.Body, err = rr.Pop(bodyLen)
	body := string(data[_heartOffset:])
	fmt.Printf("body:%v\n", body)
}

const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)



var BigEndian bigEndian

type bigEndian struct{}

func (bigEndian) Int8(b []byte) int8 { return int8(b[0]) }

func (bigEndian) PutInt8(b []byte, v int8) {
	b[0] = byte(v)
}

func (bigEndian) Int16(b []byte) int16 { return int16(b[1]) | int16(b[0])<<8 }

func (bigEndian) PutInt16(b []byte, v int16) {
	_ = b[1]
	b[0] = byte(v >> 8)
	b[1] = byte(v)
}

func (bigEndian) Int32(b []byte) int32 {
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}

func (bigEndian) PutInt32(b []byte, v int32) {
	_ = b[3]
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}