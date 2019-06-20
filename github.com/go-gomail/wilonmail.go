// https://github.com/go-gomail/gomail
// wilon mail 自定义bin，快速发送邮件方便自己使用
//
// Build：先修改config，再执行
// go build -ldflags "-s -w" -o ~/go/bin/wilonmail ./github.com/go-gomail/wilonmail.go
//
// wilonmail -h
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	gomail "gopkg.in/gomail.v2"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	to          *string = flag.String("to", "", "发送给谁，必须")
	subject     *string = flag.String("subject", "Weilong@Mac自动发送", "Use -subject <subject>")
	content     *string = flag.String("content", "", "Use -content <content>")
	contentFile *string = flag.String("contentFile", "", "Use -contentFile <contentFile>")
	attachs     arrayFlags
)

func init() {
	flag.Var(&attachs, "", "Use -attach <file> [-attach <file>...]")
}

func main() {

	flag.Parse()

	if len(*to) == 0 {
		flag.PrintDefaults()
		return
	}

	text := *content
	_, err := os.Stat(*contentFile)
	if err != nil && os.IsNotExist(err) {
	} else {
		b, err := ioutil.ReadFile(*contentFile)
		if err != nil {
			fmt.Println(contentFile, "不可读")
			return
		}
		text = text + string(b)
	}

	if len(text) == 0 {
		fmt.Println("-content 或 -contentFile 必须有一个")
		flag.PrintDefaults()
		return
	}
	htmlText := regexp.MustCompile(`\n|\r`).ReplaceAllString(text, "<br>")

	// config
	username := "13466357088@139.com"
	password := "141wilon1989"

	m := gomail.NewMessage()
	m.SetAddressHeader("From", username, "伟龙的自动提醒")
	m.SetHeader("To", *to)
	m.SetHeader("Subject", *subject)
	m.SetBody("text/html", htmlText)

	for _, v := range attachs {
		_, err := os.Stat(v)
		if err != nil && os.IsNotExist(err) {
			continue
		}
		m.Attach(v)
	}

	d := gomail.NewDialer("smtp.139.com", 25, username, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("Send Email To " + *to + " Success")
	}
}
