package main

import (
	"fmt"
	"os"

	"github.com/vcgo/kit"
	"github.com/vcgo/test"
	"github.com/xen0n/go-workwx" // package workwx
)

func main() {
	corpID := test.Config.Get("workwx.corpID").(string)
	corpSecret := test.Config.Get("workwx.corpSecret").(string)
	agentID := test.Config.Get("workwx.agentID").(int64)

	client := workwx.New(corpID)

	app := client.WithApp(corpSecret, agentID)
	// preferably do this at app initialization
	app.SpawnAccessTokenRefresher()

	// new midia
	area := kit.Area{100, 100, 400, 400}
	area = kit.Screen
	bitmap := area.Capture()
	bitmap.SavePng("./yes.png")
	area.Test("qywechat", "./tmp/")

	f, _ := os.Open("./yes.png")
	media, merr := workwx.NewMediaFromFile(f)
	kit.Fmt("..", media, merr)
	mediaRes, uerr := app.UploadTempImageMedia(media)
	kit.Fmt("upload", mediaRes, uerr)
	if uerr != nil {
		return
	}

	// send to user(s)
	to1 := workwx.Recipient{
		UserIDs: []string{"wangwl"},
	}
	res0 := app.SendTextMessage(&to1, "有新的提醒...", false)
	res1 := app.SendImageMessage(&to1, mediaRes.MediaID, false)

	fmt.Println(res0, res1)
}
