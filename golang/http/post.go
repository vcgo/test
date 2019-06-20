package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	httpDo()
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://10.80.31.133/api/test.php?233=23", strings.NewReader("account=wangwl&title=test&content=2333"))
	// req, err := http.NewRequest("POST", "https://message.ifengidc.com/api/v1/wechat", strings.NewReader("account=wangwl&title=test&content=2333"))
	if err != nil {
		// handle error
	}

	req.Header.Set("AppJWTKey", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTczNjQ4NTEsImlzcyI6Im1lc3NhZ2UtZ2F0ZXdheSIsImFwcGlkIjoibWVzc2FnZV9nTVRwQmQiLCJhcHBpZG5hbWUiOiLjgJBJVOS4reW_g-OAkSIsImFwcHR5cGUiOiIxIn0.yXxgVVI2jbjBWQjTgdBMQO1Y2plwifl_ENgHMsQHqCU")
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
