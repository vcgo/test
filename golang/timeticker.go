package main

import (
	"fmt"
	"time"
)

// 断续器
func main() {
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		n := 1
		for t := range ticker.C {
			switch true {
			case n%3 == 0:
				fmt.Println("Tick at 3", n, t)
			case n%5 == 0:
				fmt.Println("Tick at 5", n, t)
			default:
				fmt.Println("no communication +", n)
			}
			n++
		}
	}()

	go func() {
		n := 1
		for t := range ticker.C {
			switch true {
			case n%4 == 0:
				fmt.Println("Tick at 4", n, t)
			case n%6 == 0:
				fmt.Println("Tick at 6", n, t)
			default:
				fmt.Println("no communication -", n)
			}
			n++
		}
	}()

	time.Sleep(time.Second * 60)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
