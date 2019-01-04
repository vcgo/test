package main

import (
	"fmt"

	"github.com/vcgo/kit"
)

func main() {

	keyDown(1)
	kit.Sleep(555)
	keyUp(1)

}

func keyDown(i int) {
	defer keyUp(999)
	fmt.Println("keyDown ...", i)
}
func keyUp(i int) {
	fmt.Println("keyUp ...", i)
}
