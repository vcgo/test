package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	var res string
	for i := 1; i < 999999; i++ {
		res = Md5("a")
	}
	fmt.Println(res)
}

// Md5 is just md5. ^ ^
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
