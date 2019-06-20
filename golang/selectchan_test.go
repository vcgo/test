// 测试
// go test -v golang/selectchan_test.go -run Ui
package golang

import (
	"fmt"
	"testing"
	"time"
)

func TestMutiChan(t *testing.T) {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		quit <- 1
	}()

	go func() {
		for {
			select {
			case j := <-ch:
				fmt.Println(j)
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("exit")
}

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
func TestSelect(t *testing.T) {
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
