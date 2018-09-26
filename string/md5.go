package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/vcgo/kit"
)

func main() {
	kit.Fmt("Test md5", createPasswd())
}

func createPasswd() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, t.String())
	kit.Fmt("Test md5 time", t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}
