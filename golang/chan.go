package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("下载中...")
				time.Sleep(time.Second)
			}
		}
	}()
	url := "https://dldir1.qq.com/weixin/Windows/WeChat_C1018.exe"
	resp, err := http.Get(url)
	if err != nil {
	}
	fileBytes, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("config.exe", fileBytes, 0755)
	for i := 0; i < 9; i++ {
		fmt.Println("ok", i)
		if i == 5 {
			done <- true
		}
		time.Sleep(time.Second)
	}
	defer resp.Body.Close()
}

func main123() {
	roomNum := make(chan int, 1)
	go func() {
		for {
			roomNum <- rand.Intn(9)
		}
	}()
}

func mainxx() {

	ch := make(chan int, 1)

	go func() {
		inV := 1
		for {
			ch <- inV
			fmt.Println(time.Now(), "...in", inV)
			// 等待的同时会去跑下面👇
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(1)
			inV++
		}
	}()

	var outV int
	for {
		outV = <-ch
		fmt.Println(time.Now(), "...out", outV)
		for i := 0; i < 4; i++ {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println(i)
		}
	}
}
