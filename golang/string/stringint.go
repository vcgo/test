package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("string()", string(233))
	// fmt.Println("int()", int("233"))    error!
	fmt.Println("strconv.Itoa()", strconv.Itoa(233))
	n, err := strconv.Atoi("233")
	fmt.Println("strconv.Atoi()", n, err)
}
