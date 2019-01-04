package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
	"github.com/vcgo/kit"
)

func main() {
	go doSomething()
	addEvent()
}

func doSomething() {
	for i := 0; i < 9999; i++ {
		fmt.Println("...", i)
		kit.Sleep(999)
	}
}

func addEvent() {
	keve := robotgo.AddEvent("k")
	if keve == 0 {
		fmt.Println("you press...", "k")
	}
	fmt.Println("addEvent done")
	kit.Sleep(10000)
}
