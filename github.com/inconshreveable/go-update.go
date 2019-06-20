package main

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"gitee.com/libs"
	update "github.com/inconshreveable/go-update"
	"github.com/vcgo/kit"
)

// Update 之后仍运行
// 二进制文件追加几个字符串仍可以运行

// go build -o bin/selfupdate github.com/inconshreveable/go-update.go
// ./bin/selfupdate
// vim ./bin/selfupdate，已自追加字符串

func main() {
	for i := 0; i < 5; i++ {
		kit.Fmt("Update test v1", i)
		if i == 2 {
			SelfChange()
		}
		kit.Sleep(999)
	}
}

// 自动在文件尾追加随机字符串
func SelfChange() error {
	BinFullPath, _ := exec.LookPath(os.Args[0])
	println("...", BinFullPath, libs.TmpPath)
	b, _ := ioutil.ReadFile(BinFullPath)
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	str := make([]rune, 1024*5)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	var buffer bytes.Buffer
	buffer.Write(b)
	buffer.Write([]byte(string(str) + "~" + time.Now().String()))
	return update.Apply(bytes.NewReader(buffer.Bytes()), update.Options{
		OldSavePath: libs.TmpPath,
	})
}
