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

## 修改AMS OOM_ADJ
```
frameworks/base/services/core/java/com/android/server/am/ActivityManagerService.java
updateOomAdjLocked() // 方法
```


## Gradle 错误
```
Android resource linking failed
build/intermediates/incremental/mergeDebugResources/merged.dir/values-v28/values-v28.xml:7: error: resource android:attr/dialogCornerRadius not found.
build/intermediates/incremental/mergeDebugResources/merged.dir/values/values.xml:2994: error: resource android:attr/fontVariationSettings not found.
build/intermediates/incremental/mergeDebugResources/merged.dir/values/values.xml:2995: error: resource android:attr/ttcIndex not found.
error: failed linking references.
```

检查`gradle.build`的`appcompat`的版本是否相同
```
dependencies {
    implementation 'com.android.support:appcompat-v7:28.0.0-rc02'
}

```

## 查看系统默认输入法
```
adb shell settings get secure default_input_method
adb shell settings get secure enabled_input_methods
```

通过`loadStringSetting()`方法从数据库加载`default_input_method`或者`enabled_input_methods`

## 内核模块
```
insmod awusb.ko
demsg //查看模块加载失败信息
modinfo awusb.ko

```








