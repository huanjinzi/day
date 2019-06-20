# Docker

## docker build
```
docker build -t second:v1.0 .
```
## push到registry
```
docker tag busybox:latest 10.10.105.71:5000/tonybai/busybox:latest
docker push 10.10.105.71:5000/tonybai/busybox
```

## https push问题
配置文件的位置：`/etc/docker/daemon.json`
```
{ "insecure-registries":["cloud.skyworth.com:10010"] }
```
修改配置文件之后，需要重启`docker daemon`
