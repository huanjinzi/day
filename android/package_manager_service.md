# PMS

## apk
apk是zip文件

## 签名
v1 jarsigner 保护zip承载的数据
v2 apksigner 保护整个apk文件

zipalign的影响：
* v1:不破坏签名
* v2:破坏v2签名，apk降级到v1验证，一般会在v1签名中标明签名版本，防回滚攻击()

## scanDirLI
在系统开机中，通过`scanPackageChildLI()`扫描apk，如果有报错，并且：
```java
// Delete invalid userdata apps
if ((scanFlags & SCAN_AS_SYSTEM) == 0 &&
    errorCode != PackageManager.INSTALL_SUCCEEDED) {
    logCriticalInfo(Log.WARN,
        "Deleting invalid package at " + parseResult.scanFile);
        removeCodePathLI(parseResult.scanFile);
}
```