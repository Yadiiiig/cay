package core

import (
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
	s.Logger.Log("debug" + s.Lines[s.CY][:len(s.Lines[s.CY])])

	termbox.Flush()
}
