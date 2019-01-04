package main

import (
	//"encoding/json"
	"fmt"
	//"os"
	//"io/ioutil"
	"image/jpeg"
	"io"
	"log"
	"os"

	"github.com/nfnt/resize"

	// "flag"

	"image"
	"image/png"
	"strings"
)

func main() {
	imageCompress()
}

func imageCompress(
	getReadSizeFile func() (io.Reader, error),
	getDecodeFile func() (*os.File, error),
	to string, // outputPath
	Quality,
	base int, // Width
	format string) bool { // png
	/** 读取文件 */
	file_origin, err := getDecodeFile()
	defer file_origin.Close()
	if err != nil {
		fmt.Println("os.Open(file)错误")
		log.Fatal(err)
		return false
	}
	var origin image.Image
	var config image.Config
	var temp io.Reader
	/** 读取尺寸 */
	temp, err = getReadSizeFile()
	if err != nil {
		fmt.Println("os.Open(temp)")
		log.Fatal(err)
		return false
	}
	var typeImage int64
	format = strings.ToLower(format)
	/** jpg 格式 */
	if format == "jpg" || format == "jpeg" {
		typeImage = 1
		origin, err = jpeg.Decode(file_origin)
		if err != nil {
			fmt.Println("jpeg.Decode(file_origin)")
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)")
			log.Fatal(err)
			return false
		}
		config, err = jpeg.DecodeConfig(temp)
		if err != nil {
			fmt.Println("jpeg.DecodeConfig(temp)")
			return false
		}
	} else if format == "png" {
		typeImage = 0
		origin, err = png.Decode(file_origin)
		if err != nil {
			fmt.Println("png.Decode(file_origin)")
			log.Fatal(err)
			return false
		}
		temp, err = getReadSizeFile()
		if err != nil {
			fmt.Println("os.Open(temp)")
			log.Fatal(err)
			return false
		}
		config, err = png.DecodeConfig(temp)
		if err != nil {
			fmt.Println("png.DecodeConfig(temp)")
			return false
		}
	}
	/** 做等比缩放 */
	width := uint(base) /** 基准 */
	height := uint(base * config.Height / config.Width)

	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if typeImage == 0 {
		err = png.Encode(file_out, canvas)
		if err != nil {
			fmt.Println("压缩图片失败")
			return false
		}
	} else {
		err = jpeg.Encode(file_out, canvas, &jpeg.Options{Quality})
		if err != nil {
			fmt.Println("压缩图片失败")
			return false
		}
	}

	return true
}
