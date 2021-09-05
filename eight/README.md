# Go进阶训练营第8周作业

## 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能

### 命令

```
redis-benchmark -d 10 -t get,set
```

### set

|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10 |100000 requests completed in 0.95 seconds |104931.80 requests per second|
|20|    100000 requests completed in 1.00 seconds |99900.09 requests per second|
|50|   100000 requests completed in 0.97 seconds |103519.66 requests per second|
|100| 100000 requests completed in 0.92 seconds |109051.26 requests per second|
|200| 100000 requests completed in 0.98 seconds |102354.15 requests per second|
|1024| 100000 requests completed in 1.10 seconds |90579.71 requests per second|
|5120| 100000 requests completed in 1.11 seconds| 90252.70 requests per second|

### get

|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|  100000 requests completed in 0.92 seconds| 108813.92 requests per second|
|20| 100000 requests completed in 0.96 seconds| 104275.29 requests per second|
|50|  100000 requests completed in 1.00 seconds| 100401.61 requests per second|
|100 | 100000 requests completed in 0.99 seconds |100704.94 requests per second|
|200  |100000 requests completed in 1.00 seconds |99601.60 requests per second|
|1024| 100000 requests completed in 1.01 seconds| 98716.68 requests per second|
|5120 |100000 requests completed in 1.20 seconds| 83263.95 requests per second|

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

### 代码链接[github]
