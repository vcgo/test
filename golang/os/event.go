package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {

	go func() {}()
	kCh := make(chan int)
	go func() {
		for {
			fmt.Println("...event k")
			kCh <- robotgo.AddEvent("k")
		}
	}()

	sCh := make(chan int)
	go func() {
		for {
			fmt.Println("...event s")
			sCh <- robotgo.AddEvent("s")
		}
	}()

	for {
		switch {
		case <-kCh == 0:
			fmt.Println("you press...", "k")
		case <-sCh == 0:
			fmt.Println("you press...", "s")
		}
		fmt.Println("for switch")
	}
}
