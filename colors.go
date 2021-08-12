package colors

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//consts for implementations.
const prefix = "\x1b"

//output default os.Stdout
var output = os.Stdout

//Attribute is base sgr code for implementations ansi colors.
type Attribute int

// base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
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

// foreground hi-intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
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

// background hi-intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

//Colors is struct base for implements methods the colors package.
type Colors struct {
	params []Attribute
}

//NewColorsBasic is return new color instance.
func NewColors() *Colors {
	return &Colors{
		params: make([]Attribute, 0),
	}
}

// Add is used to chain SGR parameters. Use as many as parameters to combine.
func (c *Colors) Add(value ...Attribute) *Colors {
	c.params = append(c.params, value...)
	return c
}

//To256 return new colors of table the 256 colors(8-Bits).
func (c *Colors) To256(fg, bg int) *Colors {
	c.params = make([]Attribute, 0)
	c.params = append(c.params, 38, 5, Attribute(fg), 48, 5, Attribute(bg))
	return c
}

// Println formats using the default formats for its operands and writes to
// standard output.
func (c *Colors) Println(a ...interface{}) (int, error) {
	i, err := fmt.Fprintln(output, c.wrap(fmt.Sprint(a...)))
	return i, err
}

//Print formats using the default formats with not line print.
func (c *Colors) Print(a ...interface{}) (int, error) {
	i, err := fmt.Fprint(output, c.wrap(fmt.Sprint(a...)))
	return i, err
}

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Colors) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Colors) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Colors) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// sequence returns a formatted SGR sequence.
func (c *Colors) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap wraps the s string with the colors attributes.
func (c *Colors) wrap(s string) string {
	return c.format() + s + c.unformat()
}

// format return is formated text color for print.
func (c *Colors) format() string {
	return fmt.Sprintf("%s[%sm", prefix, c.sequence())
}

// unformat is responsable by reset format colors.
func (c *Colors) unformat() string {
	return fmt.Sprintf("%s[%dm", prefix, Reset)
}
