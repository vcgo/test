// 测试：如何退出一个协程 Goroutine
// 并不是中断，而是利用管道协助结束（return）
//
// go func() {
// 	for {
// 		select {
// 		case <-quit:
// 			return
// 		default:
//          // 进行主要操作，直到结束
//          fmt.Println("Goroutine...")
//      }
// }

package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Quit...")
				return
			default:
				for i := 0; i < 10; i++ {
					fmt.Println("Goroutine...", i)
					time.Sleep(time.Duration(999) * time.Millisecond)
				}
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Master...", i)
		if i == 5 {
			quit <- true
		}
		time.Sleep(time.Duration(999) * time.Millisecond)
	}
}
