## apt
````
sudo dpkg -i xx.deb //
sudo apt-get -f -y --allow-unauthenticated install

/var/cache/apt/archives //默认下载位置
/usr/share //默认安装位置
/usr/bin //可执行文件位置
/etc //配置文件位置
/usr/lib //lib文件位置
````

## Gerrit

```
ssh -p 29418 yuanhuan@192.168.1.164 gerrit ls-projects //查看Gerrit Project
ssh -p 29418 yuanhuan@192.168.1.164 gerrit review --verified 1 ad9b7af6 //code review
```
## git 命令

````
git am --skip
git am --reject --directory=src 0005-add-xrapi-and-xrcore-magic-window-impl.patch //--directory=src 需要相对目录，不能以'.','/','~'开头
git add -f 
git push origin HEAD:refs/for/master //提交代码

git instaweb --httpd=webrick -p 9998
git instaweb --httpd=webrick --stop
git config  --file ./etc/gerrit.config --unset auth.httpHeader
htpasswd -c etc/passwords admin

git commit --amend --author='yuanhuan <yuanhuan@skyworth.com>'
git remote add origin ssh://yuanhuan@192.168.1.164:29418/coocaa
git branch --set-upstream-to=origin/<branch> master

//更新子模块
git submodule init
git submodule update

git clone --recursive repository.git
````

## git多仓库管理

### repo

```
repo init -u url [options]
```
Options:

    -u: Specify a URL from which to retrieve a manifest repository. The common manifest is found at https://android.googlesource.com/platform/manifest

    -m: Select a manifest file within the repository. If no manifest name is selected, the default is default.xml.

    -b: Specify a revision, that is, a particular manifest-branch.


```
repo sync [project-list]
```

```
git fetch -all/git remote update
git rebase origin/branch
```

```
repo forall [project-list] -c command
```
Options:

    -c: Command and arguments to execute. The command is evaluated through /bin/sh and any arguments after it are passed through as shell positional parameters.

    -p: Show project headers before output of the specified command. This is achieved by binding pipes to the command's stdin, stdout, and sterr streams, and piping all output into a continuous stream that is displayed in a single pager session.

    -v: Show messages the command writes to stderr.


### gclient
工具位置
```
git clone https://chromium.googlesource.com/chromium/tools/depot_tools
```

```
gclient revinfo
```
查看git库信息
```
src: https://chromium.googlesource.com/chromium/src.git
src/buildtools: https://chromium.googlesource.com/chromium/buildtools.git@9a90d9aaadeb5e04327ed05775f45132e4b3523f
```
可以看到`@9a90d9aaadeb5e04327ed05775f45132e4b3523f`是节点信息，代码布局为`$ROOT/src`，`$ROOT/src/buildtools`


## jenkins
```
java -jar jenkins.war --httpPort=8082

JENKINS_HOME/config.xml  删除 </useSecurity> </authorizationStrategy> </securityRealm> 下的标签
```

## google source
```
https://android.googlesource.com/new-password //添加用户，避免IP限制
```

## gitiles
[gitiles]
	siteTitle = Git
	canonicalUrl = huanjinzi
	port = 8083
[markdown]
	blocknote = true
	multicolumn = true
	namedanchor = true
	smartquote = true
	toc = true

## chromium
````
gn gen --args='target_os="android"' out/Default
autoninja -C out/Default monochrome_public_apk
scp out/Default/apks/MonochromePublic.apk huanjinzi@192.168.1.113:/home/huanjinzi/
loadable_modules // 将so文件拷贝到apk
````

## adb logcat
`````
adb install -r ~/MonochromePublic.apk
adb push ~/libxrcore.so data/app/org.chromium.chrome-1/lib/arm/ || adb push ~/libxrcore.so data/app/org.chromium.chrome-2/lib/arm/
adb logcat -s chromium:* | tee log.txt
adb logcat -s TimeWarp:* chromium // logcat多标签过滤
adb shell dumpsys SurfaceFlinger
adb shell ps | grep chromium | awk '{if(NR==1) print $2}' | xargs adb shell kill //杀掉chromium浏览器
adb shell dumpsys activity service com.svr.va/.core.VAService
`````
## 文本处理与编辑
````
awk '{print $1}' //打印第一列的数据
awk '{if(NR==1) print}' //打印第一行的数据,NR行号，NF列数，
````
## 文件查找
````
find ./native_client -depth -iname .git //查找git仓库，删除加上 | xargs rm -rf
````

## ssh
````
ssh -X //图形界面
ssh-copy-id -i ~/.ssh/yuanhuan.pub ssnwt@192.168.1.172 //免密码ssh
eval "$(ssh-agent -s)" && ssh-add ~/.ssh/
````

## 系统信息
````
sudo lsb_release -a
cat /proc/version
cat /etc/issue
````

## Gradle
```
// gradle代理
systemProp.https.proxyHost=192.168.1.113
systemProp.https.proxyPort=8118
systemProp.http.proxyHost=192.168.1.113
systemProp.http.proxyPort=8118

// gradle 配置文件
~/.gradle/gradle.properties
$PROJECT_ROOT/gradle.properties

./gradlew build --refresh-dependencies //缓存刷新

// 缓存目录，出现包引入有问题的情况，可以清理缓存尝试解决
~/.gradle/caches/modules-2/files-2.1/

```

## 系统运行状态
````
htop // 进程
top // 进程
sudo iftop // 网速
````


## pdf文档
````
alias ascpdf='asciidoctor-pdf -r asciidoctor-pdf-cjk-kai_gen_gothic -a pdf-style=KaiGenGothicCN'
asciidoctor -b html5 -a icons -a toc2 -a theme=flask
:source-highlighter: pygments //注意需要空格
pip install --user pygments // pygmentize

evince //打开pdf
````

## 网络信息
```
sudo lsof -i:8080 
sudo ifconfig eno1 192.168.1.172 //修改IP，重启失效
```

## Gerrit
````
java -jar gerrit-2.16.2.war init --batch -d ./review_site //gerrit初始化

 ./review_site/bin/gerrit.sh start //gerrit启动

// 29418
ssh -p 29999 admin@192.168.1.172 gerrit //gerrit命令
ssh -p 29999 jenkins@192.168.1.172 gerrit review 2ab71c27e198f460c173f963bac34292df521cb5 --verified 1

// add maven
repositories {
    // com.hz:checker:1.0.1
    maven {
        url 'https://dl.bintray.com/huanjinzi/maven'
    }
}
````

## Clang-Format
```
clang-format -style=Chromium -i ./chrome/browser/android/vr/vr_shell_gl.h
```

## ELF
little endian 和 big endian
这两个古怪的名称来自英国作家斯威夫特的《格列佛游记》。
在该书中，小人国里爆发了内战，战争起因是人们争论，吃鸡蛋时究竟是从大头(Big-endian)敲开还是从小头(Little-endian)敲开。
为了这件事情，前后爆发了六次战争，一个皇帝送了命，另一个皇帝丢了王位。

## 网络下载
```
wget -c https://mirrors.tuna.tsinghua.edu.cn/aosp-monthly/aosp-latest.tar 

-c contine实现断点续传
-O 下载文件名字
--post-data="key=xx&value=xx"

```

## 文件系统
```
ls -i //查看inode
```

## 服务器
```
java -jar jenkins.war&
```

## bcdboot
```
https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/bcdboot-command-line-options-techref-di

bcdboot D:\Windows /s S:  /* /s system partition(EFI S:)  window系统 D:\Windows 所在盘*/
```

## diskpart
```
https://docs.microsoft.com/en-us/previous-versions/windows/it-pro/windows-server-2012-R2-and-2012/cc770877(v=ws.11)

create partition efi size=100 offset=1000
create partition msr size=200

select partition 1
select volume 1

list partition
list volume
```

## js快速网络请求
```
var url = "https://www.pgyer.com/apiv2/app/view";

var form=new FormData();

var xhr = new XMLHttpRequest();
xhr.open("POST", url , false);
xhr.send(form);
```

## curl
```
// -F表示form -X POST
curl -F 'file=@/tmp/example.ipa' -F '_api_key=5e36337b4730e0ee0fbb4bfa83242816' https://www.pgyer.com/apiv2/app/upload

-o 下载文件名字
--progress 进度
```

## keytool
```
// -keyalg RSA -keysize 2048 -validity 10000
keytool -genkey -v -alias "github" -keyalg "RSA" -keystore "huanjinzi.keystore" //创建keystore，包含一个叫github的keypair
keytool -list -keystore "huanjinzi.keystore" // 如果keystore有密码的话，需要输入密码
keytool 
```

## apk release
```
gradle assembleRelease // 编译生成apk
zipalign -v -p 4 app-unsigned.apk app-unsigned-aligned.apk // 4K对齐apk
apksigner sign --ks huanjinzi.keystore --out app-release.apk app-unsigned-aligned.apk //签名apk
apksigner verify app-release.apk //检查app是否签名

// gradle签名
signingConfigs {
        release {
            storeFile file("my-release-key.jks")
            storePassword "password"
            keyAlias "my-alias"
            keyPassword "password"
        }
    }
signingConfig signingConfigs.release
```

## Android.mk预编译apk
```
LOCAL_PATH := $(call my-dir)

# Voice Assistant
include $(CLEAR_VARS)
LOCAL_MODULE := VoiceAssistant
LOCAL_MODULE_CLASS := APPS
LOCAL_MODULE_TAGS := optional
LOCAL_BUILT_MODULE_STEM := package.apk
LOCAL_MODULE_SUFFIX := $(COMMON_ANDROID_PACKAGE_SUFFIX)
LOCAL_CERTIFICATE := platform
#LOCAL_CERTIFICATE := PRESIGNED
#LOCAL_DEX_PREOPT := false
#LOCAL_PRIVILEGED_MODULE := true
#LOCAL_OVERRIDES_PACKAGES := Home Launcher2 Launcher3
#LOCAL_MULTILIB := 32
LOCAL_SRC_FILES := $(LOCAL_MODULE).apk

include $(BUILD_PREBUILT)
```

## 压缩包预览
```
zip -v file.zip
```

## 远程桌面
```
rdesktop -f 192.168.1.129
```
windows需要关闭允许网络级别的认证


