// 阿里云 oss 官方文档
//	https://help.aliyun.com/document_detail/87712.html
//	https://github.com/aliyun/aliyun-oss-go-sdk
package main

import (
	"fmt"
	"os"

	"github.com/vcgo/test"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpoint        = ""
	accessKeyId     = ""
	accessKeySecret = ""
	bucketName      = ""
	client          oss.Client
	bucket          *oss.Bucket
)

func init() {
	endpoint = test.Config.Get("aliyunoss.endpoint").(string)
	accessKeyId = test.Config.Get("aliyunoss.accessKeyId").(string)
	accessKeySecret = test.Config.Get("aliyunoss.accessKeySecret").(string)
	bucketName = test.Config.Get("aliyunoss.bucketName").(string)

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err = client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func main() {

	// 上传本地文件。
	err := bucket.PutObjectFromFile("koala.png", "tmp/koala.png")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("...", err)
}
