* zeromq: nodePushClient(要设置大水位)-->goPullServer(也可以用zmq4或goczmq),goPushClient会丢数据(设置Send水位无效果)
* bigqueue 内存可控制,但数据文件不能自动删除,人工删除部分数据文件会报错
* dque issue和更新太少,性能太低
* badger #下载模块需要VPN,可以按前缀删除数据,然后数据文件会变少,内存会慢慢降低,支持纯内存模式,支持disk+cache模式
* goque 基于leveldb,issue和更新太少,性能可以,内存稳定占用少50M,队列出列后文件数持续增多,除非使用Drop删除队列所有文件,不能设置db文件大小
```
enqueue,dequeue,35,
enqueue,397,dequeue,134,
enqueue,492,dequeue,241,
enqueue,文件数=585,dequeue,文件数=345,内存=50M
enqueue,文件数=672,内存50M,dequeue,文件数=441,内存=51.3M
enqueue,文件数=753,内存50.3M,dequeue,文件数=532,内存=45.3M
```
* nsq 支持批量发送 每次1k*5000,每秒批量发送7次,批量大小限制5M