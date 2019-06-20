package vcgo

import (
	"fmt"
	"testing"

	"github.com/vcgo/van"
)

func TestAuthcode(t *testing.T) {
	count, key, str := 100, "2333", "weilong"
	ch := make(chan string)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			go func() {
				encodeStr, _ := van.Authcode(str, "ENCODE", key, 0)
				ch <- encodeStr
			}()
		}
	}
	m := 0
	for {
		select {
		case encodeStr := <-ch:
			decodeStr, _ := van.Authcode(encodeStr, "DECODE", key, 0)
			fmt.Println(m, decodeStr)
			m = m + 1
			if decodeStr != str {
				fmt.Println("Some Error!")
				return
			}
		}
	}
	close(ch)
}
