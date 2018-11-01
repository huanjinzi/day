## Android JNI

### Java方法签名(signature)
在JNI中，以下结构体描述Java方法与Native方法的对应关系:
```java
typedef struct {
    const char* name; //Java 方法名
    const char* signature; //Java 方法签名
    void*       fnPtr; //Native 函数指针
} JNINativeMethod;

```

```
(参数列表)[返回类型]
```

| 签名|类型 |描述|
|:-:|:-:|:--|
|`V`      |`void`   | 一般用在方法返回值 |
|`Z`      |`boolean`| |
|`B`      |`byte`   | |
|`C`      |`char`   | |
|`S`      |`short`  | |
|`I`      |`int`    | |
|`J`      |`long`   | |
|`F`      |`float`  | |
|`D`      |`double` | |
|`[`      |`数组`    | 两部分组成:[开头 结束] 开头=`[`，结束=`类型`，二维数组开头=`[[`，多维数组类推|
|`L全类名;`|`引用类型`|  三部分组成:[开头 中间 结束] 开头=`L`，中间=`全类名`，结束=`;`|


build.gradle

```groovy
android {
    
    defaultConfig {

        //(1) ndk-build 参数配置
        externalNativeBuild {
            ndkBuild {
                abiFilters 'armeabi-v7a'

                // Application.mk.
                arguments 'APP_STL=gnustl_static', // 指定 C++ 的支持水平
                        'APP_PLATFORM=android-28',
                        'NDK_TOOLCHAIN_VERSION=4.9'
                
                // cFlags '-D__STDC_FORMAT_MACROS'
                // cppFlags '-fexceptions', '-frtti'
                // targets 'libexample-one.so','my-executible-demo'
            }
        }
    }

    //(2) 指定外部编译脚本
    externalNativeBuild {
        ndkBuild {
            path 'src/main/cpp/Android.mk'
        }
    }
}
```

C++ 支持水平
|名称  | 功能|
|:-:  |:-|
|libc++ | C++17 support.|
|gnustl |Partial C++11 support. (Removed in r18.)|
|STLport |C++98 support. (Removed in r18.)|
|system |new and delete. (Deprecated in r18.)|
|none |No headers, limited C++.|

``` makefile
APP_STL=system stlport_static stlport_shared gnustl_static gnustl_shared c++_static c++_shared none

```

目录结构
```
app/
-- .externalNativeBuild/ (保存ndk-build的编译参数))
-- build/ (编译生成目录))
-- src
-- build.gradle

```

Android.mk
```makefile
LOCAL_PATH := $(call my-dir)

# module start
include $(CLEAR_VARS)

LOCAL_MODULE := native-lib

LOCAL_SRC_FILES := native-lib.cpp

# LOCAL_C_INCLUDES := android
LOCAL_LDLIBS := -llog

include $(BUILD_SHARED_LIBRARY)
# module end
```



头文件路径分析:
```
sdk/ndk-bundle/sysroot/usr/include/jni.h
sdk/ndk-bundle/sysroot/usr/include/android/log.h
sdk/ndk-bundle/sources/cxx-stl/gnu-libstdc++/4.9/include/string
```

```c++
#include <jni.h>
#include <android/log.h>
#include <string>

JNICALL
jstring get(JNIEnv *env, jobject *object) {

    jstring hello = env->NewStringUTF("hello world!");

    return hello;
}

const JNINativeMethod method[] = {
        {"get", "()Ljava/lang/String;", (void *) get}
};

#define CLASS_NAME "com/example/MainActivity"

jint JNI_OnLoad(JavaVM *vm, void *reserve) {

    __android_log_write(ANDROID_LOG_DEBUG, "TAG", "hello Logcat!");

    JNIEnv *env = NULL;

    // 这里传递 &env
    if (vm->GetEnv((void **) &env, JNI_VERSION_1_6) != JNI_OK) {
        return -1;
    }

    jclass clz = env->FindClass(CLASS_NAME);

    if (env != NULL) {
        env->RegisterNatives(clz, method, sizeof(method)/ sizeof(method[0]));
    }

    __android_log_write(ANDROID_LOG_DEBUG, "TAG", "JNI_OnLoad end");
    return JNI_VERSION_1_6;
}
```