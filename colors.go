package colors

import (
	"fmt"
)

//consts for implementations.
const (
	prefix = "\x1b"
	bright = 60
)

//colors foreground.
const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

//colors background
const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

//Colors is struct base for implements methods the colors package.
type Colors struct {
	Fg int
	Bg int
}

//NewColorsBasic is return new color instance.
func NewColors() *Colors {
	return &Colors{Fg: FgWhite, Bg: BgBlack}
}

//Println print text with colorized in terminal.
func (c *Colors) Println(content interface{}) {
	fmt.Println(c.format(content))
}

//format return of string formated.
func (c *Colors) format(content interface{}) string {
	return fmt.Sprintf("%v[%vm%v[%vm%v%v[0m", prefix, c.Fg, prefix, c.Bg, content, prefix)

}
