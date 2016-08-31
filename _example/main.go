package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/soh335/truecolor"
)

func main() {
	for h := 0; h < 360; h++ {
		tc := truecolor.New()
		tc.Add(truecolor.NewBackgrond(hsv(float64(h) / float64(360))))
		tc.Fprint(os.Stdout, "  ")
		if h%60 == 59 {
			fmt.Print("\n")
		}
	}
}

// https://ja.wikipedia.org/wiki/HSV%E8%89%B2%E7%A9%BA%E9%96%93
func hsv(h float64) color.Color {
	s := 1.0
	v := 1.0
	h *= 6.0
	i := int(h)
	f := h - float64(i)
	r, g, b := v, v, v
	switch i {
	case 0:
		g *= 1 - s*(1-f)
		b *= 1 - s
	case 1:
		r *= 1 - s*f
		b *= 1 - s
	case 2:
		r *= 1 - s
		b *= 1 - s*(1-f)
	case 3:
		r *= 1 - s
		g *= 1 - s*f
	case 4:
		r *= 1 - s*(1-f)
		g *= 1 - s
	case 5:
		g *= 1 - s
		b *= 1 - s*f
	default:
		panic("")
	}
	return &color.RGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255}
}
