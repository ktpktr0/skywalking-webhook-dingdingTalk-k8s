钉钉告警功能实践

这里使用go或者java的webhook，支持机器人的加签功能。Skywlaking 8.3版本以前需要webhook来发送到钉钉等聊天工具。以后版本内置了webhook功能，但会报错，建议使用外置webhook。
创建dingding机器人，选择加签，将token和加签值保存起来。

1、GOLANG版本的webhook
该代码依赖 github.com/weiqiang333/golang-skywalking-webhook ,增加了加签功能。

sh build/build.sh

./golang-skywalking-webhook --config configs/conf.yml


2、JAVA版本的webhook

mvn clean package

java -jar skywalking-webhook-dingding-talk.jar --spring.config.location=/tmp/application.properties
