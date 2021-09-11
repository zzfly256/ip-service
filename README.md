# IPService

纯真网络免费 IP 库解析服务。

- Go 语言编写
- 进程内缓存结果，重复的 ip 查询响应时间平均为 0.2 ms
- 内置协程数、缓存数、CPU核心数等指标上报

## 部署方法

0. 若首次部署安装，请提前准备好纯真网络的 IP 库。IP 库可以通过直接安装纯真网络客户端并将其导出为 txt 文件而获得。运行前将其命名为 `czip.txt` 并放置在 `./storage` 文件夹即可。

### 方法一：编译安装
1. 安装 golang 环境。建议 go1.13 以上。
2. 编译运行

    ```shell
    go mod download
    go get github.com/zzfly256/ip-service/src
    go build -o ./ip-service github.com/zzfly256/ip-service/src
    ./ip-service
    ```
   
### 方法二：docker

克隆本仓库后，运行一下命令
   ```shell
   docker build -t ip-service .
   docker run -d --rm -p 80:80 ip-service
   ```

## 纯真网络 ip 库获取方法
1. 打开网站：www.cz88.net 下载客户端软件，导出 txt 数据库
2. 将其转换为 UTF8 格式（采用 enca 工具）

   ```shell
   # 查看文件编码
   enca -L chinese czip.txt
   # 从 GB2312 转换为 UTF-8
   enca -L zh_CN -x UTF-8 czip.txt
   ```

## 接口清单

| 接口 | 请求方式 | 请求字段 | 说明 |
| :---- | :---- | :---- | :---- |
| /v1/query_my_ip | GET | 无 | 获取访问者自身 ip |
| /v1/query_ip_address | GET/POST | ip | 查询 ip 地址相关信息 |
| /metrics | GET | 无 | 获取程序运行相关指标(Prometheus 格式) |

详细的接口文档后续得闲再补充。

## 性能测试

1. 随机 ip 请求测试中，首次响应时长约 50ms/请求，产生缓存后约 0.2 ms/请求
2. 本人开发机压测结果如下所示

```shell
rytia@Rytia-Envy13
------------------
OS: Ubuntu 20.04.1 LTS on Windows 10 x86_64
Kernel: 4.4.0-19041-Microsoft
Uptime: 24 mins
Packages: 761 (dpkg)
Shell: zsh 5.8
Terminal: /dev/tty2
CPU: Intel i5-8250U (8) @ 1.800GHz
Memory: 6088MiB / 8038MiB

# ab -c 100 -t 10 "http://127.0.0.1:3000/v1/query_ip_address?ip=210.35.147.225"
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 5000 requests
Completed 10000 requests
Completed 15000 requests
Finished 17113 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            3000

Document Path:          /v1/query_ip_address?ip=210.35.147.225
Document Length:        114 bytes

Concurrency Level:      100
Time taken for tests:   10.041 seconds
Complete requests:      17113
Failed requests:        0
Total transferred:      4073370 bytes
HTML transferred:       1951110 bytes
Requests per second:    1704.25 [#/sec] (mean)
Time per request:       58.677 [ms] (mean)
Time per request:       0.587 [ms] (mean, across all concurrent requests)
Transfer rate:          396.15 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18  13.8     14      92
Processing:     1   40  20.8     36     197
Waiting:        1   31  17.4     27     180
Total:          1   58  22.8     53     206

Percentage of the requests served within a certain time (ms)
  50%     53
  66%     62
  75%     70
  80%     76
  90%     91
  95%    104
  98%    116
  99%    123
 100%    206 (longest request)
```