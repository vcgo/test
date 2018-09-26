package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 60; i++ {
		if i%2 == 0 {
			fmt.Println("22-", i)
		} else if i%3 == 0 {
			fmt.Println("333-", i)
		} else if i%4 == 1 {
			fmt.Println("4444-", i)
		}
	}
}
