// 作业内容：

// 1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 10
// ====== SET ======
//   100000 requests completed in 0.50 seconds
//   50 parallel clients
//   10 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 198807.16 requests per second

// ====== GET ======
//   100000 requests completed in 0.48 seconds
//   50 parallel clients
//   10 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 207039.33 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 20
// ====== SET ======
//   100000 requests completed in 0.49 seconds
//   50 parallel clients
//   20 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 206185.56 requests per second

// ====== GET ======
//   100000 requests completed in 0.51 seconds
//   50 parallel clients
//   20 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 195694.72 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 50
// ====== SET ======
//   100000 requests completed in 0.46 seconds
//   50 parallel clients
//   50 bytes payload
//   keep alive: 1

// 99.98% <= 1 milliseconds
// 100.00% <= 1 milliseconds
// 215517.25 requests per second

// ====== GET ======
//   100000 requests completed in 0.50 seconds
//   50 parallel clients
//   50 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 198807.16 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 100
// ====== SET ======
//   100000 requests completed in 0.50 seconds
//   50 parallel clients
//   100 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 200803.22 requests per second

// ====== GET ======
//   100000 requests completed in 0.51 seconds
//   50 parallel clients
//   100 bytes payload
//   keep alive: 1

// 99.97% <= 1 milliseconds
// 100.00% <= 1 milliseconds
// 194552.53 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 200
// ====== SET ======
//   100000 requests completed in 0.51 seconds
//   50 parallel clients
//   200 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 196850.39 requests per second

// ====== GET ======
//   100000 requests completed in 0.39 seconds
//   50 parallel clients
//   200 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 257069.41 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 1024
// ====== SET ======
//   100000 requests completed in 0.50 seconds
//   50 parallel clients
//   1024 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 198412.69 requests per second

// ====== GET ======
//   100000 requests completed in 0.37 seconds
//   50 parallel clients
//   1024 bytes payload
//   keep alive: 1

// 99.97% <= 1 milliseconds
// 100.00% <= 1 milliseconds
// 271739.12 requests per second


// panghaoran@DESKTOP-SK63LHQ:/bulldozer/code/jike-homework-go$ redis-benchmark -t get,set -d 5120
// ====== SET ======
//   100000 requests completed in 0.56 seconds
//   50 parallel clients
//   5120 bytes payload
//   keep alive: 1

// 100.00% <= 0 milliseconds
// 176991.16 requests per second

// ====== GET ======
//   100000 requests completed in 0.39 seconds
//   50 parallel clients
//   5120 bytes payload
//   keep alive: 1

// 99.97% <= 1 milliseconds
// 100.00% <= 1 milliseconds
// 259067.36 requests per second












// 2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
redis-benchmark -t get,set -r 10000 -q -d 10

127.0.0.1:6379> memory usage key:000000000116
(integer) 70

redis-benchmark -t get,set -r 10000 -q -d 20

127.0.0.1:6379> memory usage key:000000000116
(integer) 80

redis-benchmark -t get,set -r 10000 -q -d 100
127.0.0.1:6379> memory usage key:000000000016
(integer) 162

redis-benchmark -t get,set -r 10000 -q -d 200
127.0.0.1:6379> memory usage key:000000000016
(integer) 262

redis-benchmark -t get,set -r 10000 -q -d 1000
127.0.0.1:6379> memory usage key:000000000016
(integer) 1088

redis-benchmark -t get,set -r 10000 -q -d 5120
127.0.0.1:6379> memory usage key:000000000016
(integer) 5184