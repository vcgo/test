// github.com/go-vgo/robotgo
// 测试

package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/vcgo/kit"

	"github.com/go-vgo/robotgo"
)

type M struct{}

func main() {

	funcName := os.Args[1]

	// 动态调用
	m := reflect.ValueOf(&M{}).MethodByName(funcName)
	v := make([]reflect.Value, 0)
	m.Call(v)
}

// Event 测试按键监听
func (m *M) Event() {

	go func() {
		s := robotgo.Start()
		defer robotgo.End()
		for ev := range s {
			fmt.Println(ev)
		}
	}()

	for {
		kit.Sleep(999)
	}
}

func (m *M) KeyListen() {

	kCh := make(chan bool)
	go func() {
		for {
			fmt.Println("...event k")
			kCh <- robotgo.AddEvent("k")
		}
	}()

	sCh := make(chan bool)
	go func() {
		for {
			fmt.Println("...event s")
			sCh <- robotgo.AddEvent("s")
		}
	}()

	for {
		switch {
		case <-kCh == true:
			fmt.Println("you press...", "k")
		case <-sCh == true:
			fmt.Println("you press...", "s")
		}
		fmt.Println("for switch")
	}

	s := robotgo.Start()
	defer robotgo.End()

	for ev := range s {
		fmt.Println(ev)
	}
}
