package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/vcgo/kit"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8077", nil)
	}()
	go test1()
	test2()
}

func test1() {
	values := make([]int, 1024)
	for {
		log.Println(values)
		kit.Sleep(3333)
	}
}

func test2() {
	kit.InitLogger()
	for {
		kit.Log("...", "this", 233)
		kit.Sleep(2222)
	}
}
