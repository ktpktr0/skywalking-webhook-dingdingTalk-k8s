# golang-skywalking-webhook
skywalking alarm webhook


- install
```bash
go get -u github.com/weiqiang333/golang-skywalking-webhook
cd $GOPATH/src/github.com/weiqiang333/golang-skywalking-webhook/
bash build/build.sh
./bin/golang-skywalking-webhook help
```

- Example
```bash
./bin/golang-skywalking-webhook --config configs/conf.yml
```

# Demonstration
<img src=".static/skywalking-UI-alarm.png"/>
<p align="center">-SkyWalking alarm UI-</p>

<img src=".static/skywalking-dingding-notify.png"/>
<p align="center">-dingtalk message body-</p>
