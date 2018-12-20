package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const (
	bucket    = "dnfgo"
	accessKey = "pLXedsURF3lywR_CbVC_z1VXAuv5V3eoxkXMNPId"
	secretKey = "8fWiHXw33TlCQM38ZC9-xzGEcGAYOPLgwFkMt_bn"
)

func main() {
	key := "koala_wilon-pc-x1carbon.png"
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
}

func qiniuPut(src string, des string) {
	localFile := src
	key := des
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, key),
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "dnfgo monitor",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return
	}
}

func qiniuUpdateList() {
	bucket := "dnfgo"
	limit := 1000
	prefix := "koala_"
	delimiter := ""
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone = &storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	//初始列举marker为空
	marker := ""
	var list []storage.ListItem
	now := time.Now()
	for {
		entries, _, nextMarker, hashNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			break
		}
		for _, entry := range entries {
			if now.Unix()*10000000-entry.PutTime > 15*60*10000000 {
				continue
			}
			list = append(list, entry)
		}
		if hashNext {
			marker = nextMarker
		} else {
			break
		}
	}
	json, _ := json.Marshal(list)

	// save json
	outputFile, outputError := os.OpenFile("data.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(string(json))
	outputWriter.Flush()

	// up json
	qiniuPut("data.json", "data.json")
}
