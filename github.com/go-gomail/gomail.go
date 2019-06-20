// https://github.com/go-gomail/gomail
package main

import (
	"fmt"

	"github.com/vcgo/test"
	gomail "gopkg.in/gomail.v2"
)

func main() {

	// config
	username := test.Config.Get("email.username").(string)
	password := test.Config.Get("email.password").(string)
	dst0 := test.Config.Get("email.dst0").(string)
	// dst1 := test.Config.Get("email.dst1").(string)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", username, "伟龙的自动提醒")
	m.SetHeader("To", dst0)
	m.SetHeader("Subject", "233Hello看看!")
	m.SetBody("text/html", "233Hello <b>可靠Bob</b> and <i>Cora</i>!")
	m.Attach("./README.md")

	d := gomail.NewDialer("smtp.163.com", 25, username, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("Success", true)
	}
}
