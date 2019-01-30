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
````

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

## 服务器
````
java -jar jenkins.war&
java -jar lib/opengrok.jar -s /home/ssnwt/src/opengrok/src -d /home/ssnwt/src/opengrok/data -U http://localhost:9090/source -W /home/ssnwt/src/opengrok/config.xml -T 16 --progress -S -H -P -G -c /usr/local/bin/ctags \
-r on \
-I *.java \
-I *.c \
-I *.cc \
-I *.h \
-I *.mk \
-I *.cpp \
-I *.aidl \
-I *.sh \
-i f:*.so \
-i f:*.gz \
-i f:*.o \
-i f:*.a \
-i f:*.zip \
-i f:*.html \
-i f:*.md \
-i f:*.ninga \
-i f:*.py \
-i f:*.js \
-i d:src/abi \
-i d:src/art \
-i d:src/bionic \
-i d:src/compatibility \
-i d:src/cts \
-i d:src/dalvik \
-i d:src/developers \
-i d:src/development \
-i d:src/docs \
-i d:src/external \
-i d:src/libcore \
-i d:src/libnativehelper \
-i d:src/ndk \
-i d:src/pdk \
-i d:src/platform_testing \
-i d:src/prebuilts \
-i d:src/sdk \
-i d:src/test \
-i d:src/toolchain \
-i d:src/tools \
-i d:test \
-i d:tests
````

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

