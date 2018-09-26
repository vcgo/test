package main

import (
	"bytes"
	"crypto/sha1"
	"io/ioutil"
	"time"

	update "github.com/inconshreveable/go-update"
	"github.com/vcgo/kit"
)

// Update 之后仍运行
// 二进制文件追加几个字符串仍可以运行

// go build -o bin/selfupdate program/selfupdate.go
// ./bin/selfupdate
// vim ./selfupdate，已自追加字符串

func main() {
	for i := 0; i < 5; i++ {
		kit.Fmt("Update test v1", i)
		if i == 2 {
			b, _ := ioutil.ReadFile("selfupdate")
			h := sha1.New()
			h.Write([]byte(time.Now().String()))
			update.Apply(bytes.NewReader(h.Sum(b)), update.Options{})
		}
		kit.Sleep(999)
	}
}
