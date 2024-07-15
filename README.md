# 钉钉机器人报警模块用法

# 1.自定义钉钉机器人接入官方文档
  https://open.dingtalk.com/document/orgapp/session-management-overview

# 2.用法参考
  go get github.com/YQ131026/dingding

# 3.用法示例

```go
package main

import (
	"github.com/YQ131026/dingding"
	"log"
)

func main() {
	d := dingding.NewDingDing("text", "application/json", "https://oapi.dingtalk.com/robot/send?access_token=xxx")
	err := d.SendAlarmMessage("消息:这里是要发送警告的消息.")
	if err != nil {
		log.Fatal(err)
	}
}
```

