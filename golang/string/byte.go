package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test1 := []byte{0xAA, 0x61, 128, 64, 0x2F}

	a, b := "AA", "61"
	aa, _ := strconv.ParseInt(a, 16, 32)
	bb, _ := strconv.ParseInt(b, 16, 32)
	test2 := []byte{byte(aa), byte(bb), byte(128), 64, 0x2F}

	fmt.Println("test", test1, test2, strings.Fields("AAbbccdd", "1", 4))
}
