package test

import (
	"fmt"
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

// Config can use as: test.Config.Get()...
var (
	Config *toml.Tree
)

func Init() {
	Config, err := toml.LoadFile("config.toml")
	if err != nil {
		srcName := "config.toml.example"
		dstName := "config.toml"
		src, err := os.Open(srcName)
		if err != nil {
			return
		}
		defer src.Close()
		dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
		fmt.Println("Error ", err.Error(), Config)
	} else {
		/*

			 This is Config use func

			// retrieve data directly
			user := Config.Get("postgres.user").(string)
			password := Config.Get("postgres.password").(string)

			// or using an intermediate object
			configTree := Config.Get("postgres").(*toml.Tree)
			user = configTree.Get("user").(string)
			password = configTree.Get("password").(string)
			fmt.Println("User is", user, " and password is", password)

			// show where elements are in the file
			fmt.Printf("User position: %v\n", configTree.GetPosition("user"))
			fmt.Printf("Password position: %v\n", configTree.GetPosition("password"))

		*/
	}

}
