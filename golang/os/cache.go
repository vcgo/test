package main

import (
	"fmt"
	"os"
	"path"

	"github.com/huntsman-li/go-cache"
	_ "github.com/huntsman-li/go-cache/redis"
)

func main() {
	pwd, _ := os.Getwd()
	dir := path.Join(pwd, "caches")
	cache, err := cache.Cacher(cache.Options{
		Adapter:       "file",
		AdapterConfig: dir,
		Interval:      2,
	})
	fmt.Println("...", pwd)
	if err != nil {
		panic(err)
	}

	cache.Put("ggg", "cache", 60)
	dd := cache.Get("test")
	fmt.Println(dd)
}
