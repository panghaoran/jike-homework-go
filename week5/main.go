package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 参考 Hystrix 实现一个滑动窗口计数器
func main() {
	// 限流 计数器（固定窗口）滑动窗口 漏桶 令牌桶
	// 滑动窗口协议（Sliding Window Protocol），属于TCP协议的一种应用，用于网络数据传输时的流量控制，以避免拥塞的发生。 该协议允许发送方在停止并等待确认前发送多个数据分组。 由于发送方不必每发一个分组就停下来等待确认。 因此该协议可以加速数据的传输，提高网络吞吐量

	var size int = 100
	var reqThreshold int = 200
	var failedThreshold float64 = 0.8
	var duration time.Duration = time.Second * 2

	// 初始化
	r := NewRollingWindow(size, reqThreshold, failedThreshold, duration)
	r.Launch()     // 启动不断往里边扔桶
	r.Monitor()    // 监控
	r.ShowStatus() // 每个一秒查询当前是否处于熔断状态
}

type RollingWindow struct {
	sync.RWMutex
	broken          bool
	size            int
	buckets         []*Bucket
	reqThreshold    int       // 触发熔断的请求总数阈值
	failedThreshold float64   // 触发熔断失败率阈值
	lastBreakTime   time.Time // 上次熔断时间
	seeker          bool
	brokeTimeGap    time.Duration // 恢复时间间隔
}

func NewRollingWindow(
	size int,
	reqThreshold int,
	failedThreshold float64,
	brokeTimeGap time.Duration,
) *RollingWindow {
	return &RollingWindow{
		size:            size,
		buckets:         make([]*Bucket, 0, size),
		reqThreshold:    reqThreshold,
		failedThreshold: failedThreshold,
		brokeTimeGap:    brokeTimeGap,
	}
}

func (r *RollingWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}
}

// 获取最后一个桶
func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

func (r *RollingWindow) RecordReqResult(result bool) {
	r.GetBucket().Record(result)
}

func (r *RollingWindow) ShowAllBucket() {
	for _, v := range r.buckets {
		fmt.Printf("id: [%v] | total: [%d] | failed: [%d]\n", v.Timestamp, v.Total, v.Failed)
	}
}

func (r *RollingWindow) Launch() {
	go func() {
		for {
			r.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (r *RollingWindow) BreakJudgement() bool {
	r.RLock()
	defer r.RUnlock()
	total := 0
	failed := 0

	for _, v := range r.buckets {
		total += v.Total
		failed += v.Failed
	}

	if float64(failed)/float64(total) > r.failedThreshold && total > r.reqThreshold {
		return true
	}

	return false
}

// 监控
func (r *RollingWindow) Monitor() {
	go func() {
		for {
			if r.broken {
				if r.OverBrokenTimeGap() {
					r.Lock()
					r.broken = false
					r.Unlock()
				}
				continue
			}

			if r.BreakJudgement() {
				r.Lock()
				r.broken = true
				r.lastBreakTime = time.Now()
				r.Unlock()
			}
		}
	}()
}

func (r *RollingWindow) OverBrokenTimeGap() bool {
	return time.Since(r.lastBreakTime) > r.brokeTimeGap
}

// 每个一秒查询当前是否处于熔断状态
func (r *RollingWindow) ShowStatus() {
	go func() {
		for {
			log.Println(r.broken)
			time.Sleep(time.Second)
		}
	}()
}

// 获取熔断状态
func (r *RollingWindow) Broken() bool {
	return r.broken
}

func (r *RollingWindow) SetSeeker(status bool) {
	r.Lock()
	defer r.Unlock()
}

func (r *RollingWindow) Seeker() bool {
	return r.seeker
}

// bucket
type Bucket struct {
	sync.RWMutex
	Total     int // 请求总数
	Failed    int // 失败
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.Failed++
	}
	b.Total++
}
