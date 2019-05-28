# udp

## 广播地址
```
255.255.255.255:port
```

## 多播地址
```
224.0.0.0-239.255.255.255
```

```java
SocketAddress address = new InetSocketAddress("239.255.255.250",1900);
MulticastSocket socket = null;
try {
    socket = new MulticastSocket(address);
} catch (IOException e) {
    e.printStackTrace();
}
```
