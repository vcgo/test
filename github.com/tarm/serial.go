// U口硬件控制测试
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tarm/serial"
)

var (
	SS *serial.Port
	aa int
)

func init() {
	rand.Seed(time.Now().Unix())

	c := &serial.Config{Name: "COM3", Baud: 9600}
	SS, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err, SS)
	}
	aa = 233
}

func main() {
	fmt.Println("...", SS, aa)
	moveTo(799, 449)
}

func moveTo(x, y int) {

	for i := 0; i < 55; i++ {
		SS.Write([]byte{0xAA, 0x61, 0x04, 0x00, 127, 127, 0x00})
	}

	// 距离
	xDist, yDist := int(1600)-x, int(900)-y

	// 哪个长做master
	direction, master, slave := "x", xDist, yDist
	if xDist < yDist {
		direction = "y"
		master, slave = yDist, xDist
	}

	// 每步125
	step := 125
	masterStep, slaveStep := step, step // 初始step
	xStep, yStep := 0, 0                // 结果
	times := int(master / step)
	for i := times; i >= 0; i-- {
		// master最后一步
		if i == 0 {
			masterStep = master % step
		}
		// slave结束的早
		if slave <= slaveStep {
			slaveStep = slave
		}
		slave -= slaveStep

		if direction == "x" {
			xStep = masterStep
			yStep = slaveStep
		} else {
			xStep = slaveStep
			yStep = masterStep
		}

		SS.Write([]byte{0xAA, 0x61, 0x04, 0x00, getMoveByte(xStep), getMoveByte(yStep), 0x00})
	}
	return
}

func getMoveByte(step int) byte {
	var res int
	if step > 0 {
		res = 256 - step
	} else if step == 0 {
		res = 0
	} else {
		res = 256 + step
	}
	return byte(res)
}
