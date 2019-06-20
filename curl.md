# curl

curl  is  a tool to transfer data from or to a server, using one of the
supported protocols (DICT, FILE, FTP, FTPS, GOPHER, HTTP, HTTPS,  IMAP,
IMAPS,  LDAP,  LDAPS,  POP3,  POP3S,  RTMP, RTSP, SCP, SFTP, SMB, SMBS,
SMTP, SMTPS, TELNET and TFTP). The command is designed to work  without
user interaction.

curl offers a busload of useful tricks like proxy support, user authen‐
tication, FTP upload, HTTP post, SSL connections, cookies, file  trans‐
fer  resume,  Metalink,  and more. As you will see below, the number of
features will make your head spin!
```
// -F表示form -X POST
curl -F 'file=@/tmp/example.ipa' -F '_api_key=5e36337b4730e0ee0fbb4bfa83242816' https://www.pgyer.com/apiv2/app/upload

-o 下载文件名字
--progress 进度
```

## http请求设置header
```
curl -H "X-First-Name:Joe" http://127.0.0.1
```

设置空header
```
curl -H "X-First-Name;" http://127.0.0.1
```

## http请求response header
此时curl默认使用的http的`HEAD`方法，`-I`打印response header
```
curl -I http://127.0.0.1
```
修改curl的方法
```
curl -X GET -I http://127.0.0.1
```

## http post form data
```
curl -d "xx=ss&yy=aa" http://127.0.0.1
```

## http response
```
curl -D - -d "xx=ss&yy=aa" http://127.0.0.1
```

## multipart/form-data

```
curl -X POST -F "file=@app.apk" -F "platform=1" xx.com
```


