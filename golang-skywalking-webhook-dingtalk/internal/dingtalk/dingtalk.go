package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/weiqiang333/infra-skywalking-webhook/configs"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
[{
	"scopeId": 2,
	"name": "growing-segmentation-pid:15149@seg3",
	"id0": 47,
	"id1": 0,
	"alarmMessage": "Response time of service instance growing-segmentation-pid:15149@seg3 is more than 1000ms in 2 minutes of last 10 minutes",
	"startTime": 1568888544862
}, {
	"scopeId": 2,
	"name": "growing-segmentation-pid:11847@seg2",
	"id0": 46,
	"id1": 0,
	"alarmMessage": "Response time of service instance growing-segmentation-pid:11847@seg2 is more than 1000ms in 2 minutes of last 10 minutes",
	"startTime": 1568888544862
}]
*/
type message struct {
	ScopeId int
	Name string
	Id0	int
	Id1 int
	AlarmMessage	string
	StartTime	int
}


// Dingtalk 发送钉钉消息体
func Dingtalk(data []byte) error {

	//加签

	secret := configs.V.GetString("secret")
	timeStampNow := time.Now().UnixNano() / 1000000
	signStr :=fmt.Sprintf("%d\n%s", timeStampNow, secret)

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(signStr))
	sum := hash.Sum(nil)

	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)

	var m []message
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	contents, alertSummary := createContent(m)
	bodys := strings.NewReader(contents)
	token := configs.V.GetString("token")
	resp, err := http.Post(
		fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%v&sign=%s", token,timeStampNow,urlEncode), "application/json", bodys)
	if err != nil {
		return err
	}
	log.Println(resp.StatusCode, alertSummary)
	return nil
}


/*
状态: notify

等级: P1

告警: Skywalking
  growing-segmentation-pid:6494@seg1  id: 44  time: 1568945304861
  growing-segmentation-pid:6908@seg0  id: 43  time: 1568945304861


Item values:

0  Response time of service instance growing-segmentation-pid:6494@seg1 is more than 1000ms in 2 minutes of last 10 minutes
1  Response time of service instance growing-segmentation-pid:6908@seg0 is more than 1000ms in 2 minutes of last 10 minutes


故障修复:
*/
func createContent(message []message) (string, string) {

	contents := ""

	for _,v := range message {
		contents += fmt.Sprintf("-----来自SkyWalking的告警-----\n【time】: %v\n【scope】: %v\n【name】: %v\n【message】: %v\n\n",
			v.StartTime, v.ScopeId, v.Name, v.AlarmMessage)
	}

	data := fmt.Sprintf(`{
        "msgtype": "text",
            "text": {
            "content": "%s",
        },
        "at": {
            "isAtAll": "",
        },
    }`, contents)
	return data, contents
}

