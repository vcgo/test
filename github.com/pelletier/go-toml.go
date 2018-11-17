package main

import (
	"github.com/vcgo/kit"
	"github.com/vcgo/test"
)

func main() {
	username := test.Config.Get("email.username").(string)
	kit.Fmt("...", username)
}
