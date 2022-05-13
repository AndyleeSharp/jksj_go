## 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
### 测试机器：mbp2018, redis版本：Redis 6.2.7 (e6f67092/0) 64 bit   
### redis-benchmark -d 10 -t get,set
   ```
   ====== SET ======
   100000 requests completed in 1.66 seconds
   50 parallel clients
   10 bytes payload
   Summary:
  throughput summary: 60204.70 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.426     0.200     0.375     0.703     0.943     1.631
  ====== GET ======
  100000 requests completed in 1.64 seconds
  50 parallel clients
  10 bytes payload
  Summary:
  throughput summary: 60827.25 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.421     0.192     0.367     0.711     0.943     1.791
```
        

### redis-benchmark -d 20 -t get,set
```
====== SET ======
  100000 requests completed in 1.58 seconds
  50 parallel clients
  20 bytes payload
Summary:
  throughput summary: 63171.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.406     0.160     0.375     0.591     0.815     3.375
====== GET ======
  100000 requests completed in 1.83 seconds
  50 parallel clients
  20 bytes payload
  Summary:
  throughput summary: 54674.69 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.470     0.232     0.407     0.839     1.183     2.727
```
### redis-benchmark -d 50 -t get,set
```
  ====== SET ======
  100000 requests completed in 1.60 seconds
  50 parallel clients
  50 bytes payload
  Summary:
  throughput summary: 62383.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.415     0.200     0.367     0.671     0.911     7.407
  ====== GET ======
  100000 requests completed in 1.55 seconds
  50 parallel clients
  50 bytes payload
  Summary:
  throughput summary: 64391.50 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.398     0.216     0.375     0.583     0.735     1.399
  ```
### redis-benchmark -d 100 -t get,set
```
  ====== SET ======
  100000 requests completed in 1.52 seconds
  50 parallel clients
  100 bytes payload
  Summary:
  throughput summary: 65919.58 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.390     0.160     0.367     0.527     0.735     1.935
  ====== GET ======
  100000 requests completed in 1.44 seconds
  50 parallel clients
  100 bytes payload
  Summary:
  throughput summary: 69396.25 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.369     0.168     0.359     0.447     0.607     1.087
  ```
### redis-benchmark -d 200 -t get,set
```
====== SET ======
  100000 requests completed in 1.65 seconds
  50 parallel clients
  200 bytes payload
Summary:
  throughput summary: 60459.49 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.425     0.192     0.383     0.671     0.847     1.991
====== GET ======
  100000 requests completed in 1.54 seconds
  50 parallel clients
  200 bytes payload
Summary:
  throughput summary: 64766.84 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.395     0.176     0.367     0.575     0.783     1.215
```
### redis-benchmark -d 1000 -t get,set
```
====== SET ======
  100000 requests completed in 1.57 seconds
  50 parallel clients
  1000 bytes payload
Summary:
  throughput summary: 63897.76 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.402     0.176     0.375     0.607     0.807     1.839
====== GET ======
  100000 requests completed in 1.56 seconds
  50 parallel clients
  1000 bytes payload
Summary:
  throughput summary: 64143.68 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.399     0.200     0.383     0.527     0.751     1.287
```
### redis-benchmark -d 5000 -t get,set
```
====== SET ======
  100000 requests completed in 1.65 seconds
  50 parallel clients
  5000 bytes payload
Summary:
  throughput summary: 60606.06 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.424     0.168     0.383     0.647     0.879     2.335
====== GET ======
  100000 requests completed in 1.70 seconds
  50 parallel clients
  5000 bytes payload
Summary:
  throughput summary: 58823.53 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.435     0.240     0.407     0.647     0.855     1.311

```

### 小结
总体来看性能还是比较稳定，不同value大小的情况下，set和get的性能都能达到5w/s,当value比较大时有略微的下降。经测试当value达到20k时，性能开始显著下降。