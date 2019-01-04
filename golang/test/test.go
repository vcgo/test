package main

import (
	"io"
	"os"
	"strconv"
)

func main() {
	for i := 0; i < 122; i++ {
		var src *os.File
		src, _ = os.Open("a.png")
		defer src.Close()
		dst, _ := os.OpenFile(strconv.Itoa(i)+".png", os.O_WRONLY|os.O_CREATE, 0644)
		defer dst.Close()
		io.Copy(dst, src)
	}
}
