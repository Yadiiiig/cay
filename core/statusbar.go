package core

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

var prevWidth struct {
    cursorIndicator int
}

// Initializes the status bar and draws the info for the first time
func (s *State) LoadBar() {

	width, height := termbox.Size()

	status := "Hello status bar"
	for i := 0; i < len(status); i++ {
		termbox.SetChar(i, height-1, rune(status[i]))
	}

	text := "[" + strconv.Itoa(s.CY) + ":" + strconv.Itoa(s.CX) + "]"
	lenght := width - len(text)
	for i := width - 1; i > lenght-1; i-- {
		termbox.SetChar(i, height-1, rune(text[i-lenght]))
    }
    prevWidth.cursorIndicator = lenght
}

// Updates the data on the bar
func (s *State) UpdateBar() {

	width, height := termbox.Size()

	status := "Status Hello"
	for i := 0; i < len(status); i++ {
		termbox.SetChar(i, height-1, rune(status[i]))
	}

	text := "[" + strconv.Itoa(s.CY) + ":" + strconv.Itoa(s.CX) + "]"
	lenght := width - len(text)
    if lenght > prevWidth.cursorIndicator {
        for i := width-1; i > prevWidth.cursorIndicator-1; i-- {
            termbox.SetChar(i, height-1, rune(0))
        }
    }
	for i := width - 1; i > lenght-1; i-- {
		termbox.SetChar(i, height-1, rune(text[i-lenght]))
    }
    prevWidth.cursorIndicator = lenght
}
