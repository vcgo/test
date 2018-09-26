package main

import (
	"fmt"
	"strconv"

	"github.com/vcgo/kit"
)

func main() {
	for i := 0; i < 9; i++ {
		go down("a" + strconv.Itoa(i))
		kit.Sleep(999)
	}
}

func down(k string) {
	fmt.Println("down", k)
	defer up(k)
}
func up(k string) {
	fmt.Println("up", k)
}
