package main

import (
	"fmt"

	"github.com/vcgo/test/interfacetest"
)

// interface 使用

// 定义新 type 类型
type student int

// 该类型若适应接口，必须定义接口的所有方法
func (t student) Say() {
	fmt.Println("...interfacetest Say")
}

func (t student) Sing() {
	fmt.Println("...interfacetest Sing")
}

// 接口使用
func main() {
	// 定义接口变量
	var tt interfacetest.Human
	// 变量赋值为定义的 type
	tt = student(23)
	// 使用方法
	tt.Say()
}
