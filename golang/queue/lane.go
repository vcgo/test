package main

import (
	"strconv"
	"time"

	"github.com/vcgo/kit"

	. "gopkg.in/oleiade/lane.v1"
)

var PriorityQueue *PQueue

func init() {
	PriorityQueue = NewPQueue(MINPQ)
}

func main() {

	// 生产者：每秒增加一个
	go func() {
		i := 0
		for {
			PriorityQueue.Push("oqlzUvga_cC6gaOD0I"+strconv.Itoa(i), i)
			i++
			// kit.Sleep(9)
		}
	}()

	// 消费者：每30s消费掉
	go func() {
		tck := time.NewTicker(10 * time.Second)
		defer tck.Stop()
		for {
			select {
			case <-tck.C:
				for {
					value, _ := PriorityQueue.Pop()
					if value == nil {
						break
					}
					kit.Fmt("...", value)
				}
			}
		}
	}()

	// master wait
	for i := 0; ; i++ {
		kit.Fmt("wait", i)
		kit.Sleep(1000)
	}

}
