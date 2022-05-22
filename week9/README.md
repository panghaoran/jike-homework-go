1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

定长 简单和底层 MTU
基于分隔符 要结合转义符实现 HTTP 数据量小 几byte 合适
header 定长 、body 不定长 Dubbo
header 不定长、body 也不定长 goIM

2. 实现一个从 socket connection 中解码出 goim 协议的解码器
main.go