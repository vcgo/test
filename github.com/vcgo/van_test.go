package vcgo

import (
	"fmt"
	"testing"
	"time"

	"github.com/vcgo/van"
)

func TestToken(t *testing.T) {
	key := "github/wilon_weilong"
	str := "abcdefghijklmno"
	encodeStr, _ := van.Authcode(str, "ENCODE", key, 3)
	fmt.Println(time.Now(), encodeStr)
}

// go test -v github.com/vcgo/van_test.go -run TestAuthcode
func TestAuthcode(t *testing.T) {
	count, key, str := 2, "2333", "weilong"
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
			fmt.Println(m, encodeStr, decodeStr)
			m = m + 1
			if decodeStr != str {
				fmt.Println("Some Error!")
				return
			}
		}
	}
	close(ch)
}
