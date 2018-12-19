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
## git 命令

````
git am --skip
git am --reject --directory=src 0005-add-xrapi-and-xrcore-magic-window-impl.patch //--directory=src 需要相对目录，不能以'.','/','~'开头
git add -f 
git push origin HEAD:refs/for/master //提交代码
````

## chromium
````
gn gen --args='target_os="android"' out/Default
autoninja -C out/Default monochrome_public_apk
````

## adb logcat
`````
adb install -r MonochromePublic.apk
adb push ~/libxrcore.so data/app/org.chromium.chrome-1/lib/arm/ || adb push ~/libxrcore.so data/app/org.chromium.chrome-2/lib/arm/
adb logcat -s chromium:*
adb logcat -s TimeWarp:* chromium // logcat多标签过滤
adb shell dumpsys SurfaceFlinger
adb shell ps | grep chromium | awk '{if(NR==1) print $2}' | xargs adb shell kill //杀掉chromium浏览器
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
scp out/Default/apks/MonochromePublic.apk huanjinzi@192.168.1.113:/home/huanjinzi/
````

## 系统信息
````
sudo lsb_release -a
cat /proc/version
cat /etc/issue
````

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

