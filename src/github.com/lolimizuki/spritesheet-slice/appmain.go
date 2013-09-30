package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strconv"
)

func main() {
	imageFile, fileErr := os.Open("case1.png")
	defer imageFile.Close()
	if fileErr != nil {
		panic(fileErr)
	}

	// originImage, format, imgErr := image.Decode(bufio.NewReader(imageFile))
	originImage, _, imgErr := image.Decode(bufio.NewReader(imageFile))
	if imgErr != nil {
		panic(imgErr)
	}

	tempImage := image.NewNRGBA(originImage.Bounds())

	for w := 0; w < originImage.Bounds().Size().X; w++ {
		for h := 0; h < originImage.Bounds().Size().Y; h++ {
			fmt.Println(originImage.At(w, h))
		}
	}

	// tempImage.SetNRGBA(x, y, c)

	save, _ := os.Create("new.png")
	png.Encode(bufio.NewWriter(save), tempImage)

	// fmt.Println(format)

	// rect := originImage.Bounds()
	// fmt.Println("Load image size=(", rect.Size().X, ",", rect.Size().Y, ") ", "format =", format)

	// var numSliceInWidth, numSliceInHeight int
	// numSliceInWidth = 5
	// numSliceInHeight = 2

	// // widthPerSlice := rect.Size().X / numSliceInWidth
	// // heightPerSlice := rect.Size().Y / numSliceInHeight

	// for w := 0; w < numSliceInWidth; w++ {
	// 	for h := 0; h < numSliceInHeight; h++ {
	// 		sliceImage := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{10, 10}})
	// 		saveFileName := "new" + strconv.Itoa(w) + "_" + strconv.Itoa(h) + "." + format
	// 		saveImgFile, _ := os.Create(saveFileName)
	// 		png.Encode(saveImgFile, sliceImage)
	// 	}
	// }
}

func temp() {
	var rr io.Reader
	rr.Read(nil)

	image.Decode(nil)

	fmt.Println("aaa")

	png.Encode(nil, nil)

	fmt.Println(strconv.Itoa(123))

	fmt.Println(color.NRGBA{32, 32, 32, 128})
}
