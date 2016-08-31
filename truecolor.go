package truecolor

import (
	"fmt"
	"image/color"
	"io"
	"strconv"
	"strings"
)

type Parameter interface {
	String() string
}

const escape = "\x1b"

// sgr code

type Code int

const (
	Reset Code = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ImageNegative
	Conceal
	CrossedOut

	Foreground = 38
	Background = 48
)

func (c Code) String() string {
	return strconv.Itoa(int(c))
}

type Color struct {
	color color.Color
	code  Code
}

func NewForeground(color color.Color) *Color {
	return &Color{
		color: color,
		code:  Foreground,
	}
}

func NewBackgrond(color color.Color) *Color {
	return &Color{
		color: color,
		code:  Background,
	}
}

func (c *Color) String() string {
	r, g, b, _ := c.color.RGBA()
	return fmt.Sprintf("%d;2;%d;%d;%d", c.code, r>>8, g>>8, b>>8)
}

type TrueColor struct {
	parameters []Parameter
}

func New() *TrueColor {
	return &TrueColor{
		parameters: []Parameter{},
	}
}

func (t *TrueColor) Add(p Parameter) {
	t.parameters = append(t.parameters, p)
}

func (t *TrueColor) Reset() {
	t.parameters = []Parameter{}
}

func (t *TrueColor) set(w io.Writer) {
	args := make([]string, len(t.parameters))
	for i, p := range t.parameters {
		args[i] = p.String()
	}
	fmt.Fprintf(w, "%s[%sm", escape, strings.Join(args, ";"))
}

func (t *TrueColor) reset(w io.Writer) {
	fmt.Fprintf(w, "%s[%dm", escape, Reset)
}

func (t *TrueColor) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	t.set(w)
	defer t.reset(w)
	return fmt.Fprint(w, a...)
}

func (t *TrueColor) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	t.set(w)
	defer t.reset(w)
	return fmt.Fprintf(w, format, a...)
}

func (t *TrueColor) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	t.set(w)
	defer t.reset(w)
	return fmt.Fprintln(w, a...)
}
