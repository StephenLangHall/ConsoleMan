package main

import (
	"fmt"
	"github.com/kaey/framebuffer"
)

func devidable (a int, b int) bool {
	return a % b == 0
}

func main () {
	fb, err := framebuffer.Init("/dev/fb0")
	if err != nil {
		fmt.Print(err)
		fmt.Print("Cannot open fb")
	}
	defer fb.Close()

	var w,h int = fb.Size()

	if (len(Image) / 6) % 2 != 0 {
		fmt.Print("Image length is not multiple of 6")
	}

	var Image2 []int
	for i,e := range Image {
		if devidable(i,1) {
			if Image[i] < 0 {
				Image2 = append(Image2, h - Image[i])
			} else {
				Image2 = append(Image2, Image[i])
			}
		} else if devidable(i,2) {
			if Image[i] < 0 {
				Image2 = append(Image2, w - Image[i])
			} else {
				Image2 = append(Image2, Image[i])
			}
		} else {
			Image2 = append(Image2, Image[i])
		}
		if e < 0 {
			Image2 = append(Image2, w-e)
		}
	}

	for i := 0; i < len(Image2) / 6; i++ {
		fb.Clear(0,0,0,0)
		fb.WritePixel(Image[i],Image[i+1],Image[i+2],Image[i+3],Image[i+4],Image[i+5])
	}
}
