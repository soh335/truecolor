package truecolor

import (
	"bytes"
	"image/color"
	"testing"
)

func TestFprint(t *testing.T) {
	var b bytes.Buffer
	tc := New()
	tc.Add(NewBackgrond(color.RGBA{255, 0, 0, 255}))
	tc.Add(NewForeground(color.RGBA{100, 100, 100, 255}))
	tc.Add(Italic)
	tc.Fprintf(&b, "hoge")
	if e, g := "\x1b[48;2;255;0;0;38;2;100;100;100;3mhoge\x1b[0m", b.String(); e != g {
		t.Errorf("expected %s but got %s", e, g)
	}
}
