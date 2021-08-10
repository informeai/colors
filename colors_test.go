package colors

import (
	"fmt"
	"log"
	"testing"
)

//go test -v -run ^TestNewColors
func TestNewColors(t *testing.T) {
	c := NewColors()
	if c == nil {
		t.Errorf("TestNewColors(): got -> %v, want: Color{}", c)
	}
	text := fmt.Sprintf("\x1b[%vm\x1b[%vm ola \x1b[0m", c.Fg, c.Bg)
	log.Println(text)
}

//go test -v -run ^TestPrintln
func TestPrintln(t *testing.T) {
	c := Colors{Fg: FgRed, Bg: BgBlack}
	c.Println("red")
}
