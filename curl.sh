#!/usr/bin/env bash

Host="http://localhost:8090"

# create source
function PostRecommend() {
    curl \
    -X POST \
    -F "index=1" \
    -F "state=1" \
    -F "channel=1" \
    ${Host}/console/cms/recommend -v
}

# create source
function PostSource() {
    curl \
    -X POST \
    -F "name=东邪西毒2" \
    -F "url=http://test.com/test.mp4" \
    -F "resolution=5" \
    -F "type=4" \
    -F "video=1" \
    ${Host}/console/cms/source -v
}

# update source
function PutSource() {
    curl \
    -X PUT \
    -F "url=update_http://test.com/test.mp4" \
    -F "type=1" \
    ${Host}/console/cms/source/${1} -v
}

# delete source
function DeleteSource() {
    curl \
    -X DELETE \
    ${Host}/console/cms/source/${1} -v
}

# create video
function PostVideo() {
    curl \
    -X POST \
    -F "name=东邪西毒2" \
    -F "source=0" \
    -F "year=1994" \
    -F "channel=1" \
    -F "title=威尼斯国际电影节最佳摄影 " \
    -F "post=https://gss2.bdstatic.com/9fo3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=4ef0c3e95edf8db1a8237436684ab631/728da9773912b31ba2c10f588618367adab4e116.jpg" \
    -F "description=欧阳锋(张国荣饰)在兄长成婚的那天离家出走，因为大嫂(张曼玉饰)是他最爱的女人。他隐居在沙漠小镇，经营着一家旅店，做着为他人寻找杀手的生意。风流剑客黄药师(梁家辉饰)是欧阳锋的好友，每年都要与锋畅饮一次。他既迷恋着好友的老婆桃花(刘嘉玲饰)，也暗恋欧阳锋的大嫂。慕容燕(林青霞饰)因为黄药师酒后的一句话深深爱上了黄药师，却因为得不到爱人的真心而伤心欲绝，她将自己幻想成慕容嫣请求欧阳锋杀死黄药师，却始终下不了狠心。" \
    -F "director=王家卫" \
    -F "actor=张国荣，林青霞，梁家辉，张曼玉，梁朝伟，刘嘉玲，张学友，杨采妮" \
    -F "area=香港" \
    -F "price=0" \
    -F "score=9.0" \
    -F "duration=6000" \
    -F "resolution=5" \
    -F "live=false" \
    -F "publish=false" \
    -F "mark=false" \
    ${Host}/console/cms/video -v
}

# create video
function PutVideo() {
    curl \
    -X PUT \
    -F "channel=1" \
    -F "title=update_威尼斯国际电影节最佳摄影_update" \
    -F "post=update_https://gss2.bdstatic.com" \
    -F "description=update_欧阳锋(张国荣饰)在兄长成婚的那天离家出走，因为大嫂(张曼玉饰)是他最爱的女人。他隐居在沙漠小镇，经营着一家旅店，做着为他人寻找杀手的生意。风流剑客黄药师(梁家辉饰)是欧阳锋的好友，每年都要与锋畅饮一次。他既迷恋着好友的老婆桃花(刘嘉玲饰)，也暗恋欧阳锋的大嫂。慕容燕(林青霞饰)因为黄药师酒后的一句话深深爱上了黄药师，却因为得不到爱人的真心而伤心欲绝，她将自己幻想成慕容嫣请求欧阳锋杀死黄药师，却始终下不了狠心。" \
    -F "director=update_王家卫" \
    -F "actor=update_张国荣，林青霞，梁家辉，张曼玉，梁朝伟，刘嘉玲，张学友，杨采妮" \
    -F "area=update_香港" \
    -F "price=1" \
    -F "score=10" \
    -F "duration=1000" \
    -F "resolution=1" \
    -F "live=true" \
    -F "publish=true" \
    -F "mark=true" \
    ${Host}/console/cms/video/${1} -v
}

# get video
function GetVideo() {
    url="${Host}/console/cms/video"
    if [[ "${1}X" != "X" ]] ; then
        url=${url}/${1}
    fi
    curl \
    -X GET \
    ${url} -v -o result.json
}

function DeleteVideo() {
    url="${Host}/console/cms/video/${1}"
    curl \
    -X DELETE \
    ${url} -v
}


function Upload() {
    curl \
    -X POST \
    -F "file=@bilibili.apk" \
    -F "platform=1" \
    ${Host}/console/appStore/Upload -v
}

function Publish() {
    curl \
    -X POST \
    -F "category=3" \
    -F "Developer=skyworth" \
    -F "Description=a new app" \
    -F "Cover=cover.com" \
    -F "ScreenShot=[screen1.shot,screen2.shot]" \
    ${Host}/console/appStore/Publish -v
}

function GetAppList() {
    curl \
    -X POST \
    -F "platform=1" \
    ${Host}/console/appStore/GetAppList -v
}

function Usage() {
    echo "Usage:"
    echo "    curl interface [params..]"
}

if [[ $1 = "-h" ]]; then
    Usage
    exit 0
fi

pwd
$1 $2
