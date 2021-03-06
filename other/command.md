## 环境变量
```
$USER //用户名
$? //上次命令执行的结果
```

## apt
````
sudo dpkg -i xx.deb //
sudo apt-get -f -y --allow-unauthenticated install

// -y --yes, --assume-yes Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively.
// -f --fix-broken Fix; attempt to correct a system with broken dependencies in place.

/var/cache/apt/archives //默认下载位置
/usr/share //默认安装位置
/usr/bin //可执行文件位置
/etc //配置文件位置
/usr/lib //lib文件位置
locate /sources.list // 查看source.list

// source.list的存放位置 
/etc/apt/sources.list
/etc/apt/sources.list.d/

````

## source-list
sources.list 文件中的条目一般都有如下所示的形式：
```
deb http://site.example.com/debian distribution component1 component2 component3
deb-src http://site.example.com/debian distribution component1 component2 component3
```
### 档案类型 (Archive type)
条目的第一个词 `deb` 或是 `deb-src` 表明了所获取的软件包档案类型。

其中，

*. deb 档案类型为二进制预编译软件包，一般我们所用的档案类型。
*. deb-src 档案类型为用于编译二进制软件包的源代码。

### 仓库地址 (Repository URL)
条目的第二个词则是软件包所在仓库的地址。我们可以更换仓库地址为其他地理位置更靠近自己的镜像来提高下载速度。

常用镜像地址列表：

```
Debian
https://www.debian.org/mirror/list
Ubuntu
http://wiki.ubuntu.org.cn/源列表
```

###发行版 (Distribution)
跟在仓库地址后的是发行版。发行版有两种分类方法，一类是发行版的具体代号，如 `xenial`,`trusty`, `precise` 等；
还有一类则是发行版的发行类型，如`oldstable`, `stable`, `testing` 和 `unstable`。

另外，在发行版后还可能有进一步的指定，如 `xenial-updates`, `trusty-security`, `stable-backports` 等。

###软件包分类 (Component)
跟在发行版之后的就是软件包的具体分类了，可以有一个或多个。

不同的 Linux 发行版对软件有着不同的分类，如：

####Debian
#####main
包含符合 DFSG 指导原则的自由软件包，而且这些软件包不依赖不符合该指导原则的软件包。这些软件包被视为 Debian 发型版的一部分。
#####contrib
包含符合 DFSG 指导原则的自由软件包，不过这些软件包依赖不在 main 分类中的软件包。
#####non-free
包含不符合 DFSG 指导原则的非自由软件包。
####Ubuntu
#####main
官方支持的自由软件。
#####restricted
官方支持的非完全自由的软件。
#####universe
社区维护的自由软件。
#####multiverse
非自由软件。


## ln
```
ln -s /home/huanjinzi/go/src/cms/cms-backend /home/huanjinzi/cms

1.确保 /home/huanjinzi/cms 不存在
2.上面的命令执行后的效果
cd /home/huanjinzi/go/src/cms/cms-backend 与 cd /home/huanjinzi/cms 有相同的效果
```


## bash
```
bash -c "pwd"
```

## Gerrit

```
ssh -p 29418 yuanhuan@192.168.1.164 gerrit ls-projects //查看Gerrit Project
ssh -p 29418 yuanhuan@192.168.1.164 gerrit review --verified 1 ad9b7af6 //code review
// 创建账户
cat ~/.ssh/id_watcher.pub | ssh -p 29418 review.example.com gerrit create-account --group "'Non-Interactive Users'" --ssh-key - watcher
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

git rebase -i HEAD~4 // fixup squash将会合并commit

git checkout -b branch_name v1.0 //基于Tag创建分支
git push origin :tag_name //push tag 到远程仓库，删除远程分支
git tag -d tag_name // 删除tag 

//创建新分支
git checkout -b|B 
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
gn ls out/Default/ device/vr
autoninja -C out/Default device/vr:vr
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
//文件浏览
wc -l //查看文件行数
head -n10
tail -n+10
tail -n10
tail -n+10 | head //第10行到第20行
less more//浏览

column -t -s ":"
sed s/^/chromium_/ 在行首插入
sed s/$/chromium_/ 在行尾插入

sed "s/^/chromium\//" 插入特殊字符

sed a\chromium_ 在行首插入新行
sed a\chromium_ 在行尾插入新行

awk '{print $1}' //打印第一列的数据
awk '{if(NR==1) print}' //打印第一行的数据,NR行号，NF列数，
find ./s8000/ -name "*.tgz" | awk '{if( NR%2 == 1) print  }' | xargs sudo rm //等间隔删除文件
````
## find
````
find ./native_client -depth -iname .git //查找git仓库，删除加上 | xargs rm -rf
find ./ -maxdepth 1 -name default_volume_tables.xml //在当前目录查找
````

## ls
```
ls -l //查看文件列表

```

##解压
```
tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```

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

## gerrit
````
java -jar gerrit-2.16.2.war init --batch -d ./review_site //gerrit初始化

 ./review_site/bin/gerrit.sh start //gerrit启动

// 29418
ssh -p 29999 admin@192.168.1.172 gerrit //gerrit命令
ssh -p 29999 jenkins@192.168.1.172 gerrit review 2ab71c27e198f460c173f963bac34292df521cb5 --verified 1
ssh -p 29999 admin@192.168.1.172 delete-project delete Chromium --yes-really-delete --force  //删除项目
cat ~/.gclient_entries | grep src/ | column -t -s \' | awk '{print $1}' | column -t -s : | awk '{print $1}' | sed "s/^/ chromium\//" | xargs bash -c

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
--no-check-certificate //关闭证书验证
-O- -q //输出到屏幕 -qO-

wget https://192.168.1.113:10443/api/getCategoryList --post-data="apiKey=321324dsfdsfs1233sasd" --no-check-certificate

```

## 文件系统
```
ls -i //查看inode
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

## unzip
```
unzip -l file.zip //查看文件列表
unzip VoiceAssistant.apk META-INF/CERT.RSA 解压指定文件

zip -r foo.zip META-INF/ // 压缩META-INF/ 文件夹到foo.zip
```

## keytool
```
// -keyalg RSA -keysize 2048 -validity 10000
keytool -genkey -v -alias "github" -keyalg "RSA" -keystore "huanjinzi.keystore" //创建keystore，包含一个叫github的keypair
keytool -list -keystore "huanjinzi.keystore" // 如果keystore有密码的话，需要输入密码
keytool -export -alias github -file test.crt -keystore huanjinzi.keystore //提取证书

keytool -printcert -rfc -file CERT.RSA
openssl pkcs7 -inform DER -in META-INF/CERT.RSA -noout -print_certs -text


//keystore没法直接获取私钥，需要转换成 PKCS12再获取私钥
keytool -importkeystore -srckeystore keystore.jks -destkeystore keystore.p12 -deststoretype PKCS12 //keystore类型转换
openssl pkcs12 -in keystore.p12 -nokeys -out my_key_store.crt //导出证书
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

### LOCAL_EXPORT_C_INCLUDES
让依赖这个项目的模块可以找到此模块中的`header files`，即暴露此模块的`header files`，
> 注意：此变量需要使用绝对路径

### LOCAL_C_INCLUDES
本模块依赖的`header files`
> 注意：此变量需要使用绝对路径比较保险

## 压缩包预览
```
zip -v file.zip
```

## 远程桌面
```
rdesktop -f 192.168.1.129
```
windows需要关闭允许网络级别的认证

## Android Studio
没有语法提示

. 菜单`File`->`Invalidate Caches/Restart...`
. 省电模式


## NDK
```
APP_STL := c++_static
APP_LDFLAGS := -L/home/huanjinzi/workspace/project/8895A71/out/target/product/hmd8895/obj_arm/lib //链接动态库
```

## Libreoffice
```
sudo add-apt-repository ppa:libreoffice/ppa
sudo apt-get update && sudo apt-get -y dist-upgrade
sudo apt-get install libreoffice

// dist-upgrade in addition to performing the function of upgrade, 
// also intelligently handles changing dependencies with new versions of packages;
```

## 进程
```
pstree //进程树查看
pstree -apl [pid] //查看进程线程
ps -Lf 915 //查看进程线程

```

## centos防火墙
```
netstat -ntpl | grep 3306

firewall-cmd --state
systemctl start firewalld.service
firewall-cmd --list-ports
firewall-cmd --permanent --zone=public --add-port=8080/tcp
firewall-cmd --reload

iptables -I INPUT -p tcp -m state --state NEW -m tcp --dport 3306 -j ACCEPT
iptables -L -n
iptables -D INPUT -p tcp -m state --state NEW -m tcp --dport 3306 -j ACCEPT
```

## 创建空文件
```
dd if=/dev/zero of=test_file bs=100M count=1
```

## 文件
```
file [file]

xxd -l 3 [file]
xxd -s +2 -l 2

watch //时间间隔执行命令

```

## 正则表达式
```
grep -ve [[:digit:]] // -v反向匹配
```

## cron log
```
sudo vim /etc/rsyslog.d/50-default.conf
cron.*   /var/log/cron.log #将cron前面的注释符去掉
sudo service rsyslog restart
tail -f /var/log/cron.log

crontab -l | crontab -
*/1 * * * * huanjinzi ~/cms/run.sh
```

## 字体
```
fc-list :lang=zh //查看中文字体

```

## 大小写转换
```
echo abc | tr a-z A-Z
```

##正则表达式

### 数字版本号匹配
例如：`1.0.1`,`1.12`
```
^[0-9]+([.][0-9]+)+$
```

## 显卡信息查看
```
lspci
```

## covert命令
convert命令可以用来转换图像的格式，支持JPG, BMP, PCX, GIF, PNG, TIFF, XPM和XWD等类型，下面举几个例子: 
```
  convert  xxx.jpg  xxx.png   将jpeg转成png文件 
  convert  xxx.gif   xxx.bmp  将gif转换成bmp图像 
  convert  xxx.tiff    xxx.pcx   将tiff转换成pcx图像 
  ```
还可以改变图像的大小: 
```
  convert -resize 1024x768  xxx.jpg   xxx1.jpg    将图像的像素改为1024*768，注意1024与768之间是小写字母x 
  convert -sample 50%x50%  xxx.jpg  xxx1.jpg   将图像的缩减为原来的50%*50% 
```
旋转图像：
``` 
convert -rotate 270 sky.jpg sky-final.jpg      将图像顺时针旋转270度 
```
使用-draw选项还可以在图像里面添加文字： 
```
convert -fill black -pointsize 60 -font helvetica -draw 'text 10,80 "Hello, World!"'  hello.jpg  helloworld.jpg
```
在图像的10,80 位置采用60磅的全黑Helvetica字体写上 Hello, World! 
```

1、批量图像格式转换

如果想将某目录下的所有jpg文件转换为png文件，只要在命令行模式下输入: 
```
　　for %f in (*.jpg) do convert "%f" "%~nf.png"
```

2、对所有图像进行同一操作

譬如，批量生成某目录下所有PNG图像文件的缩略图(大小为80×40): 
```
　　for %f in (*.png) do convert "%f" -sample 80×40 "%~nf_sample.png"
```
类似的，将某目录下所有PNG图像旋转90度的操作为： 
```
　　for %f in (*.png) do convert "%f" -rotate 90 "%~nf_rotate.png" 
```
还可以进行批量裁剪、淡化、抖动、炭化、加边框、圆角等等一系列操作

3、在图像上加上文字说明 

如果你有大量图片需要发布，在所有图片上加上版权说明是很明智的做法。
```
　　convert channel.png -fill red -pointsize 13 -draw 'text 10,15 "lifesinger 2006"' 2.png
```

## OTA
```
./build/tools/releasetools/ota_from_target_files -i old-ota-middle.zip new-ota-middle.zip ~/update.zip
```


## IP设置
```
    ifconfig eth0 128.224.156.81 up  //一般的嵌入式linux中设置IP。
    ifconfig eth0 128.224.156.81 netmask 255.255.255.0 up
    netcfg eth0 up dhcp   //通过dhcp 自动获取ip，网关，掩码等。
```

## gateway
```
    route add default gw 128.224.156.1  //一般的嵌入式linux会用这个命令添加网关,但是android中不用。直接“netcfg eth0 up dhcp” 自动分配IP等。
```

## dns
```
echo "nameserver 128.224.160.11" > resolv.conf  //一般嵌入式linux中设置dns的方法
nameserver 128.224.160.11

setprop net.dns1 128.224.160.11   //android中设置DNS的命令。
setprop net.dns2 147.11.100.30
```
## 自动获取IP
```
./dhcpcd -U eth1
broadcast_address='192.168.100.255'
dhcp_lease_time='30'
dhcp_message_type='5'
dhcp_server_identifier='192.168.100.1'
ip_address='192.168.100.10'
network_number='192.168.100.0'
routers='192.168.100.1'
subnet_cidr='24'
subnet_mask='255.255.255.0'
```
 

## 自动获取IP的详细调用过程
```
./dhcpcd -h eth1 -d
```

## 查看内核编译选项
```
cat /boot/config-`uname -r`
```

## cgroup
挂载cgroup
```
mkdir /mnt/cgroup
mount -t cgroup cgroup /mnt/cgroup/
```
ubuntu默认挂载目录`/sys/fs/cgroup`

## mkdir
```
mkdir {A,B} //创建多级目录
```

## mount
```
cat /etc/mtab //查看系统中已经挂载的设备
```

## 分区
```
mkfs.ext4 /dev/vdc

```

## ffmpeg
```
ffmpeg -re -i 4k.mp4 -c:v copy -c:a copy -f flv -y rtmp://127.0.0.1/live/test
find /home/storage/backup/jinding/517/2019-05-16/ -name "*.flv" | xargs -I @ ./ffmpeg -re  -i @ -c:v copy -c:a copy -f flv -y rtmp://127.0.0.1/live/517
ffmpeg -i input.flv output.mp4
ffmpeg -i rtmp://server/live/originalStream -c:a copy -c:v libx264 -vpre slow -f flv rtmp://server/live/h264Stream
ffmpeg -i input.avi -b:v 64k -bufsize 64k output.mp4
```

## android pm
```
cmd package list packages
pm path com.android.chrome
```

## all winner
```
#include <linux/sched.h>
#include <linux/sched/signal.h>
#include <linux/hardirq.h>     //<linux/smp_lock.h>
```

## linux mod
```
insmod xx.ko
modinfo xx.ko
lsmod
```

## cmd package
```
cmd package resolve-activity -a android.intent.action.MAIN -c android.intent.category.LAUNCHER com.android.chrome
```

## am
```
am start com.android.chrome com.android.chrome/com.google.android.apps.chrome.Main
```

## jar
创建jar包
```
jar cf class.jar com/
```

## javac
```
javac -cp ~/workspace/sdk/platforms/android-28/android.jar com/ssnwt/http/*
```

## find
删除cpp文件
```
find  /home/huanjinzi/workspace/project/VRLib/DMR/src/main/jni/include -regex "\(.*[.]cpp\)" | xargs rm
```

## so
```
ldd xx.so //查看so库的依赖
readelf -h xx.so //查看so库的header信息

//动态库连接配置，lib*.so*
ldconfig -v
/etc/ld.so.conf.d/

// 环境变量，指定so库的位置
LD_LIBRARY_PATH=[path]
```
一般配置优先级，编译参数 > 环境变量 > 自定义 > 系统默认

1. 每个动态库有一个包含了真正的库代码的文件名，通常被称为库的 `realname` ，命名格式通常为:
`libxxx.so.x.y.z`，其中so后缀中的`x`为主版本号，`y`为副版本号，`z`为发行版本号。例如，我的linux系统机器上zlib共享库的`realname`为 `libz.so.1.2.8`，这个文件是含有可执行的二进制代码的。

2. 每个动态库都有一个以`lib`为前缀且以`.so.x`为结尾的被称为`soname`:
的特定名称，其中`x`为主版本号，`soname`命名格式通常为`libxxx.so.x`。例如，我的linux系统机器上`zlib`共享库的`soname`为`libz.so.1`。这个`soname`包含了动态库的主版本号，这个`soname`一般会包含在库代码的头文件中，这个可以使用 `readelf -d` 读取出来，使用这个动态库的程序的二进制ELF的头文件中包含有这个动态库的`soname`。程序运行时会按照这个名称去找真正的库文件。

3. 此外，编译链依赖了共享库的应用模块时，链接器只认不带任何版本号的共享库名， 可以将库名称作`linker name`:
例如，我的linux系统机器上`zlib`共享库的`linkername`为`libz.so`。也即，链接使用了动态库的程序时查找的动态库名称。例如：`gcc -o test test.o -lz` , 链接时就会找`libz.so`。若没有这个文件，链接器就报错。

## smb
```
smbclient //192.168.1.172/ssnwt
```

## rpm转deb
```
sudo alien xx.rpm
sudo dpkg -i xxxx.deb #安装 
```


