package main

import (
	"encoding/json"
	"fmt"

	jpush "github.com/DeanThompson/jpush-api-go-client"
	"github.com/DeanThompson/jpush-api-go-client/push"
	"github.com/vcgo/test"
)

var (
	appKey       string
	masterSecret string
)

func main() {

	appKey := test.Config.Get("jpush.appKey").(string)
	masterSecret := test.Config.Get("jpush.masterSecret").(string)

	// platform 平台
	platform := push.NewPlatform()
	platform.All()
	// audience 对象，表示消息受众
	audience := push.NewAudience()
	// audience.SetTag([]string{"weilong", "weicong"})
	audience.All()
	// notification 对象，表示 通知，传递 alert 属性初始化
	notification := push.NewNotification("Notification")
	// android 平台专有的 notification，用 alert 属性初始化
	androidNotification := push.NewAndroidNotification("已经超时！！！https://github.com/DeanThompson/jpush-api-go-client/blob/master/jpush_test.go")
	androidNotification.Title = "掉线了！！"
	androidNotification.AddExtra("keykeykeykeykey", "valuevaluevaluevaluevalue")
	// iOS 平台专有的 notification，用 alert 属性初始化
	// iosNotification := push.NewIosNotification("iOS Notification Alert")
	// iosNotification.Badge = 1
	notification.Android = androidNotification
	// notification.Ios = iosNotification

	// message 对象，表示 透传消息，用 content 属性初始化
	message := push.NewMessage("Message Content must not be empty")
	message.Title = "Message Title"
	options := push.NewOptions()

	payload := push.NewPushObject()
	payload.Platform = platform
	payload.Audience = audience
	payload.Notification = notification
	payload.Message = message
	payload.Options = options

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("res", err)
	}
	fmt.Println("payload:", string(data), "\n")

	// Push 会推送到客户端
	jclient := jpush.NewJPushClient(appKey, masterSecret)
	result, err := jclient.Push(payload)

	fmt.Println("res", result)
	fmt.Println("err", err)
}
