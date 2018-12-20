package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:58888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//获取服务端消息
	go ioCopy(os.Stdout, conn)
	//将用户输入的文本消息发送到到服务端
	ioCopy(conn, os.Stdin)
}

func ioCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
