package main

import (
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

//func ScreenShot(displayIndex int) (*image.RGBA, error)
// save *image.RGBA to filePath with PNG format.

func save(img *image.RGBA, filePath string) {
	name := time.Unix(int64(time.Now().Unix()), 0).Format("2006_01_02__15-04-05")
	filePath = filePath + name + ".png"
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	err = png.Encode(file, img)
	if err != nil {
		return
	}
}

func main() {
	//使用 GetDisplayBounds获取指定屏幕显示范围，全屏截图
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	save(img, "d:/Error_Picture/")
}
