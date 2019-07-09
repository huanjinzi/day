# APK

## keystore
a keystore file contructure.
```
cert_1/
    public/
    private/
cert_2/
...
cert_n
```

1.generate keystore and rsa key.
```
keytool -genkey -v -alias "github" -keyalg "RSA" -keysize 2048 -keystore "huanjinzi.ks"
```
2.list certs.
```
keytool -list -keystore "huanjinzi.ks"
```

3.export one cert.
```
keytool -export -alias github -file pub.crt -keystore huanjinzi.ks
```

4.cat cert info.
```
keytool -printcert -file pub.crt
```

5.cat private key and cert info.
```
openssl pkcs12 -in huanjinzi.ks
```

## Zip
zip file structure.
```
file_1
file_2
...
file_n
center directory
```

1.list file
```
unzip -l file.zip
```

2.extract a file.
```
unzip file.zip ${path}
```

3.compress a directory.
```
zip -r file-zip ${dir}
```

4.add new file to `file.zip`.
```
zip file.zip ./*
zip file.zip new_file
```

## ApkSigner
1.cat apk cert.
```
apksigner verify --print-certs app.apk
```

2.sign apk(no matter apk has siged).
```
apksigner sign --ks release.jks --in app.apk --out app-signed.apk
```
