package core

import (
	"github.com/nsf/termbox-go"
)

// InputCapture captures all the input from mouse/keyboard
// It will only ever return true if the application should exit
func (s *State) InputCapture(ev *termbox.Event) bool {
	if ev.Type == termbox.EventKey {
		if ev.Ch != 0 {
			s.KeyStrokeMap(ev.Ch)
		} else if ev.Key == termbox.KeyEsc {
			return true
		} else if ev.Key == termbox.KeySpace {
			s.AddSpace()
		} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
			s.BackSpace()
		} else if ev.Key == termbox.KeyEnter {
			s.NewLine()
		} else {
			switch ev.Key {
			case termbox.KeyArrowLeft:
				s.MoveLeft()
			case termbox.KeyArrowRight:
				s.MoveRight()
			case termbox.KeyArrowUp:
				s.MoveUp()
			case termbox.KeyArrowDown:
				s.MoveDown()
			}
		}
	}

	return false

	// tvw.Box.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {

	// 	// This checks if it's a 'typeable key' is
	// 	if e.Rune() != 0 {
	// 		KeyStroke(tvw, e.Rune())
	// 		return nil
	// 	}

	// 	return nil
	// })
}

func foo(s string) {

}
