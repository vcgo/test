package main

import (
	"crypto/rc4"
	"flag"
	"fmt"
)

// 
func main() {
	var key, str string
	flag.StringVar(&key, "key", "123abc", "Secret key.")
	flag.StringVar(&str, "str", "abc123", "Source string.")

	var byteKey []byte = []byte(key)        //初始化用于加密的KEY
	rc4obj1, _ := rc4.NewCipher(byteKey)    //返回 Cipher
	rc4str1 := []byte(str)                  //需要加密的字符串
	plaintext := make([]byte, len(rc4str1)) //
	rc4obj1.XORKeyStream(plaintext, rc4str1)
	//XORKeyStream方法将src的数据与秘钥生成的伪随机位流取XOR并写入dst。dst和src可指向同一内存地址；但如果指向不同则其底层内存                不可重叠
	// plaintext就是你加密的返回过来的结果了，注意：plaintext则为 base-16 编码的字符串，每个字节使用 2 个字符表示 必须格式化成字符串

	stringinf1 := fmt.Sprintf("%x\n", plaintext) //转换字符串
	fmt.Println(stringinf1)
}
