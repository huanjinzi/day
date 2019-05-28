# tcp

## 三次握手
```
client(SYN_SENT)      ---- SYN --->    server(LISTEN)
client(ESTABLISH)     <--- ACK ----    server(SYN_RCVD)
client(ESTABLISH)     ---- ACK --->    server(ESTABLISH)
```

## 四次挥手
```
client(FIN_WAIT1)     ---- FIN --->    server(CLOSE_WAIT)
client(FIN_WAIT2)     <--- ACK ----    server(CLOSE_WAIT)
client(TIME_WAIT)     <--- FIN ----    server(LAST_ACK)
client(TIME_WAIT)     ---- ACK --->    server(LAST_ACK)
```

`TIME_WAIT`的意义：让连接双方的状态正常进入关闭状态，主要是`server`端可能会超时重传`FIN`，所以需要等待`2ML`。

1.可靠地实现TCP全双工连接的终止

为了保证A发送的最后一个ACK报文段能够到达B。

A给B发送的ACK可能会丢失，B收不到A发送的确认，B会超时重传FIN+ACK报文段，此时A处于2MSL时间内，就可以收到B重传的FIN+ACK报文段，接着A重传一次确认，重启2MSL计时器。最后，A和B都能够正常进入到CLOSED状态。

如果A在发完ACK后直接立即释放连接，而不等待一段时间，就无法收到B重传的FIN+ACK报文段，也就不会再次发送确认报文段，这样，B就无法按照正常步骤进入CLOSED状态。

2.允许旧的报文段在网络中消逝  

A发送确认后，该确认报文段可能因为路由器异常在网络中发生“迷途”，并没有到达B，该确认报文段可以称为旧的报文段。A在超时后进行重传， 发送新的报文段，B在收到新的报文段后进入CLOSED状态。在这之后，发生迷途的旧报文段可能到达了B，通常情况下，该报文段会被丢弃，不会造成任何的影响。但是如果两个相同主机A和B之间又建立了一个具有相同端口号的新连接，那么旧的报文段可能会被看成是新连接的报文段，如果旧的报文段中数据的任何序列号恰恰在新连接的当前接收窗口中，数据就会被重新接收，对连接造成破坏。为了避免这种情况，TCP不允许处于TIME_WAIT状态的连接启动一个新的连接，因为TIME_WAIT状态持续2MSL，就可以保证当再次成功建立一个TCP连接的时，来自之前连接的旧的报文段已经在网络中消逝，不会再出现在新的连接中，(这是因为每个TCP Packet 都有TTL)。

## Time-To-Live
TTL(Time-To-Live)的作用是限制数据包在网络中存在的时间，防止数据包不断的在IP互联网络上循环,TTL指定数据包被路由器丢弃之前允许通过的最大网段数量，是IP数据包在网络中可以转发的最大跳数(跃点数)，TTL位于IPv4包的第9个字节，是一个8 bit字段。

TTL字段由数据包的发送者设置，路由器转发数据包时，至少将TTL减小1。路由器将会丢弃TTL=0的数据包，并向数据包源地址发送一个类型11的ICMP报文，表示time exceeded（TTL为0），由发送者决定是否要重发。
