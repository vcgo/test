package main

import (
	"fmt"

	"github.com/vcgo/kit"
)

func main() {
	a := fmt.Sprintf("%x", "aa")
	b := fmt.Sprintf("%010d", 10)
	kit.Fmt("...", a, b, kit.Base64Encode("AA=="))
	kit.Fmt("...", kit.Base64Decode("gw=="))
}
