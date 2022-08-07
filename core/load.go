package core

import "github.com/nsf/termbox-go"

func (s *State) LoadLine() {
	for i := 0; i < len(s.Lines[s.CY]); i++ {
		termbox.SetChar(s.CX, s.CY, rune(s.Lines[s.CY][i]))
	}
	termbox.Flush()
}
