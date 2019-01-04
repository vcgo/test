package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/vcgo/kit"
)

// Renameself when the binary run completion.
func main() {
	fullPath, _ := exec.LookPath(os.Args[0])
	dirPath := filepath.Dir(fullPath)
	kit.Fmt("Test", fullPath, dirPath)
	return
	err := os.Rename(fileName, md5Now())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func md5Now() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}
