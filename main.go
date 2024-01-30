package main

import (
	"fmt"
	"github.com/kaey/framebuffer"
)

func main () {
	fb, err := framebuffer.Init("/dev/fb0")
	if err != nil {
		fmt.Print(err)
		fmt.Print("Cannot open fb")
	}
	defer fb.Close()

	if (len(Image) / 6) % 2 != 0 {
		fmt.Print("Image length is not multiple of 6")
	}

	for i := 0; i < len(Image) / 6; i++ {
		fb.Clear(0,0,0,0)
		fb.WritePixel(i,i+1,i+2,i+3,i+4,i+5)
	}
}
