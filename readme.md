* zeromq: nodePushClient(要设置大水位)-->goPullServer(也可以用zmq4或goczmq),goPushClient会丢数据(设置Send水位无效果)
* bigqueue 内存可控制,但数据文件不能自动删除,人工删除部分数据文件会报错
* dque issue和更新太少,性能太低
* badger #下载模块需要VPN,可以按前缀删除数据,然后数据文件会变少,内存会慢慢降低,支持纯内存模式,支持disk+cache模式
* goque 基于leveldb,issue和更新太少,性能可以,内存稳定占用少50M,队列出列后文件数持续增多,除非使用Drop删除队列所有文件,不能设置db文件大小
* pogreb issue和更新太少,性能可以,内存太高,数据文件不支持删除
* goleveldb 性能高,低内存,数据压缩率高,支持按key范围来释放数据文件
* nutsdb(待测试性能)

```
enqueue,dequeue,35,
enqueue,397,dequeue,134,
enqueue,492,dequeue,241,
enqueue,文件数=585,dequeue,文件数=345,内存=50M
enqueue,文件数=672,内存50M,dequeue,文件数=441,内存=51.3M
enqueue,文件数=753,内存50.3M,dequeue,文件数=532,内存=45.3M
```
* nsq (go-nsq)
  * 支持批量发送 每次1k*5000,每秒批量发送7次,批量大小限制5M,接收消息1W/s
  * 接收后数据文件自动缩小,
  * 内存中保留1W条消息
  * 每个channel(队列)各保存一份数据
  * 自带集群方案
  ```
  #!/bin/bash
  #集群服务
  tmux new -AdDEP -s nsqlookupd01 ./nsqlookupd
  tmux new -AdDEP -s nsqlookupd02 ./nsqlookupd -tcp-address 0.0.0.0:4162 -http-address 0.0.0.0:4163
  #队列服务
  tmux new -AdDEP -s nsqd01 ./nsqd --lookupd-tcp-address=127.0.0.1:4160 --lookupd-tcp-address=127.0.0.1:4162
  tmux new -AdDEP -s nsqd02 ./nsqd --lookupd-tcp-address=127.0.0.1:4160 -tcp-address=0.0.0.0:4152 -http-address=0.0.0.0:4153 -data-path=/tmp --lookupd-tcp-address=127.0.0.1:4162
  #web控制台
  tmux new -AdDEP -s nsqadmin01 ./nsqadmin --lookupd-http-address=127.0.0.1:4161 --lookupd-http-address=127.0.0.1:4163
  ```
  ```
  #!/bin/bash
  #集群服务
  tmux kill-session -t nsqlookupd01
  tmux kill-session -t nsqlookupd02
  #队列服务
  tmux kill-session -t nsqd01
  tmux kill-session -t nsqd02
  #web控制台
  tmux kill-session -t nsqadmin01
  ```
* nats-streaming (默认--store MEMORY) 
  * 支持持久订阅
  `./nats-streaming-server` 内存模式
  * 同步发布消息(Hello World)较慢<1W/s,异步发送较快~18W/s, 异步发送1K数据~12W/s
  * 客户端内存稳定很小~10M,接收很快~34W/s
  * 服务端内存较高>300M~2.4G 
  `./nats-streaming-server --store FILE --dir /tmp/nats` 文件模式
  * 客户端:异步写入1K数据~2.6W/s,读取~5W/s,内存~20M
  * 服务端:内存稳定~250M,150W数据文件大小1.1G,需要指定(--file_read_buffer_size 1024),否则读取时内存会很高甚至GB
  * 消息超出指定个数时,自动覆盖老数据
