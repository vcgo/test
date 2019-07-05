package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"runtime"

	"github.com/go-vgo/robotgo"

	"github.com/disintegration/imaging"
)

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 1. 按BitmapStr读取图片
	whereBitmap := robotgo.BitmapStr("b79,66,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")
	bb := robotgo.ToBitmapBytes(whereBitmap)

	// 2. 转换成 image.Image
	img, _, _ := image.Decode(bytes.NewReader(bb))
	// 2.1 测试下图片是否匹配
	imgn, err := imaging.Open("./tmp/test.bmp")
	fmt.Println("imaging...", img.Bounds().String(), imgn.Bounds().String())

	// 3. 图片切割成几小块
	rects := []image.Rectangle{}
	startX, startY := 8, 8
	leftRightLen, upDownLen := 20, 15
	sqrLen := 15
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			x0 := startX + i*leftRightLen
			y0 := startY + j*upDownLen
			x1 := x0 + sqrLen
			y1 := y0 + sqrLen
			rects = append(rects, image.Rect(x0, y0, x1, y1))
		}
	}

	// 4. 测试小块
	for _, r := range rects {
		rectcropimg := imaging.Crop(img, r)
		fileName := "./tmp/T_" + r.String() + ".png"
		imaging.Save(rectcropimg, fileName)
		findBitmap := robotgo.OpenBitmap(fileName)
		x, y := robotgo.FindBitmap(findBitmap, whereBitmap, 0.1)
		fmt.Println("RES:", x, y, r)
	}
	return

	// test image.NRGBA 转 string
	rectcropimg := imaging.Crop(img, image.Rect(9, 9, 19, 19))
	buf := new(bytes.Buffer)
	err = png.Encode(buf, rectcropimg)
	newBytes := buf.Bytes()
	pngStr := string(newBytes)
	fmt.Println("imaging...", pngStr)
	// findBmt := robotgo.BitmapStr(kit.Base64Encode(string(newBytes)))
	// x, y := robotgo.FindBitmap(bitmap, findBmt, 0.14)
	// fmt.Println("imaging...", x, y, string(newBytes))

	// test
	bit := robotgo.OpenBitmap("./tmp/centercrop.png")
	robotgo.TostringBitmap(bit)
	fmt.Println("imaging...", robotgo.ToBitmap(bit))

	// save cropped image
	err = imaging.Save(rectcropimg, "./tmp/test_a.png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// everything ok
	fmt.Println("Done")

}
