# Android APK签名
```
apksigner verify --print-certs app.apk //查看app签名
apksigner sign --ks release.jks --in app.apk --out app-signed.apk
```
