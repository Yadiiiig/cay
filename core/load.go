package core

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func (s *State) LoadLine() {
	w, _ := termbox.Size()
	//s.Logger.Log("Width:" + fmt.Sprint(w))
	for i := 0; i < w; i++ {
		termbox.SetChar(i, s.CY, 0)
	}
	//termbox.Flush()
	for i := 0; i < len(s.Lines[s.CY]); i++ {
		termbox.SetChar(i, s.CY, rune(s.Lines[s.CY][i]))
	}

	termbox.Flush()
}

func (s *State) LoadIndexLine(length, index, y int) {
	s.Logger.Log(fmt.Sprintf("len %d, index %d, y %d", length, index, y))
	for i := index; i <= length; i++ {
		termbox.SetChar(i, y, 0)
	}

	termbox.Flush()
}
