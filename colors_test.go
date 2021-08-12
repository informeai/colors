package colors

import (
	"log"
	"strings"
	"testing"
)

//go test -v -run ^TestNewColors
func TestNewColors(t *testing.T) {
	c := NewColors()
	if c == nil {
		t.Errorf("TestNewColors(): got -> %v, want: Color{}", c)
	}
}

//go test -v -run ^TestAdd
func TestAdd(t *testing.T) {
	c := NewColors()
	c = c.Add(FgGreen)
	if len(c.params) == 0 {
		t.Errorf("TestAdd(): got -> %v, want: len > 0", len(c.params))
	}
}

//go test -v -run ^TestSequence
func TestSequence(t *testing.T) {
	c := NewColors()
	c = c.Add(FgGreen, Bold)
	s := c.sequence()
	if !strings.Contains(s, ";") && len(c.params) > 1 {
		t.Errorf("TestSequence(): string bad formated.")
	}
}

//go test -v -run ^TestWrap
func TestWrap(t *testing.T) {
	c := NewColors()
	c = c.Add(FgMagenta, Bold)
	s := c.wrap("magenta")
	if len(s) == 0 && len(c.params) > 1 {
		t.Errorf("TestWrap(): string bad formated.")
	}
	log.Println(s)
}

//go test -v -run ^TestFormat
func TestFormat(t *testing.T) {
	c := NewColors()
	c = c.Add(FgMagenta)
	s := c.format()
	if !strings.HasPrefix(s, "\x1b") {
		t.Errorf("TestFormat(): string not contains prefix.")
	}
	log.Println(s)
}

//go test -v -run ^TestUnformat
func TestUnformat(t *testing.T) {
	c := NewColors()
	c = c.Add(FgMagenta)
	s := c.unformat()
	if !strings.HasPrefix(s, "\x1b") || !strings.HasSuffix(s, "\x1b[0m") {
		t.Errorf("TestUnformat(): string not contains prefix or suffix correct.")
	}
}

//go test -v -run ^TestPrintln
func TestPrintln(t *testing.T) {
	c := NewColors()
	c = c.Add(FgGreen, Bold)
	_, err := c.Println("bold green.")
	if err != nil {
		t.Errorf("TestPrintln(): got -> %v,want: nil", err)
	}

}

//go test -v -run ^TestPrint
func TestPrint(t *testing.T) {
	c := NewColors()
	c = c.Add(FgRed, Bold)
	_, err := c.Print("bold green.")
	if err != nil {
		t.Errorf("TestPrint(): got -> %v,want: nil", err)
	}

}

//go test -v -run ^TestSprint
func TestSprint(t *testing.T) {
	c := NewColors()
	c = c.Add(FgRed)
	s := c.Sprint("green.")
	if len(s) == 0 {
		t.Errorf("TestSprint(): got -> %v,want: len > 0", len(s))
	}

}

//go test -v -run ^TestSprintln
func TestSprintln(t *testing.T) {
	c := NewColors()
	c = c.Add(FgYellow)
	s := c.Sprintln("yellow.")
	if len(s) == 0 {
		t.Errorf("TestSprintln(): got -> %v,want: len > 0", len(s))
	}

}

//go test -v -run ^TestSprintf
func TestSprintf(t *testing.T) {
	c := NewColors()
	c = c.Add(FgCyan)
	s := c.Sprintf("%v", "cyan.")
	if len(s) == 0 {
		t.Errorf("TestSprintln(): got -> %v,want: len > 0", len(s))
	}

}

//go test -v -run ^TestTo256
func TestTo256(t *testing.T) {
	c := NewColors()
	c = c.To256(92, 15)
	if c == nil {
		t.Errorf("TestTo256(): got -> %v, want: != nil", c)
	}
	s := c.Sprint("this is 8-Bit colors.")
	log.Println(s)

}

//go test -v -run ^TestToRGB
func TestToRGB(t *testing.T) {
	c := NewColors()
	foreground := RGB{R: 123, G: 159, B: 23}
	background := RGB{R: 0, G: 0, B: 0}
	c = c.ToRGB(foreground, background)
	if c == nil {
		t.Errorf("TestRGB(): got -> %v, want: != nil", c)
	}
	s := c.Sprint("this is RGB colors.")
	log.Println(s)

}
