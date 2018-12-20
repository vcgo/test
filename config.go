package test

import (
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

var Config *toml.Tree

// Config can use as:
// 		test.Conf()
// 		test.Config.Get()...
func init() {
	// 注意局部变量和全局变量
	conf, err := toml.LoadFile("config.toml")
	Config = conf
	if err != nil {
		srcName := "config.toml.example"
		dstName := "config.toml"
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
		defer src.Close()
		dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
		fmt.Println("Error ", err.Error())
		return
	} else {
		return
	}
}
