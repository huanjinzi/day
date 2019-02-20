# Android

## 音量曲线调节
DEFAULT_DEVICE_CATEGORY_SPEAKER_VOLUME_CURVE
```
etc/default_volume_tables.xml
system/etc/default_volume_tables.xml
frameworks/av/services/audiopolicy/config/default_volume_tables.xml
```
验证命令:
```
adb pull etc/default_volume_tables.xml .
...
adb remount
adb push default_volume_tables.xml system/etc/default_volume_tables.xml

adb shell stop
adb shell start

```

## 系统状态
```
dumpsys cpuinfo //进程cpu使用信息
```







