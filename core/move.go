package core

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func (s *State) MoveLeft() {
	if s.CX != 0 {
		s.CX -= 1
	} else {
		if s.CY != 0 {
			s.CY -= 1
			s.CX = len(s.Lines[s.CY])
		} else {
			return
		}
	}

	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()
	// s.Logger.Log(fmt.Sprintf("x: %d, y: %d | length line: %d | current char: %s rune: %d", s.CX, s.CY, len(s.Lines[s.CY]), string(termbox.GetCell(s.CX, s.CY).Ch), termbox.GetCell(s.CX, s.CY).Ch))
}

func (s *State) MoveRight() {
	if len(s.Lines[s.CY])-1 != s.CX && len(s.Lines[s.CY])-1 > s.CX {
		s.CX += 1
	} else if s.CX == 0 && len(s.Lines[s.CY]) == 1 || s.CX == len(s.Lines[s.CY])-1 {
		s.CX += 1
	} else {
		if len(s.Lines)-1 != s.CY {
			s.CY += 1
			s.CX = 0
		} else {
			return
		}
	}
	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()
	s.Logger.Log(fmt.Sprintf("x: %d, y: %d | length line: %d | current char: %s rune: %d", s.CX, s.CY, len(s.Lines[s.CY]), string(termbox.GetCell(s.CX, s.CY).Ch), termbox.GetCell(s.CX, s.CY).Ch))
}

/*
Moving up and done result in the cursor ending up at the front or back from a line.
To Handle this, we have to come up with some calculation that will put
*/

func (s *State) MoveUp() {
	if s.CY != 0 {
		s.CY -= 1
		if s.CX > len(s.Lines[s.CY])-1 {
			if len(s.Lines[s.CY]) != 0 {
				s.CX = len(s.Lines[s.CY]) - 1
			} else {
				s.CX = 0
			}
		}
	} else {
		return
	}

	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()
	s.Logger.Log(fmt.Sprintf("x: %d, y: %d | length line: %d | current char: %s rune: %d", s.CX, s.CY, len(s.Lines[s.CY]), string(termbox.GetCell(s.CX, s.CY).Ch), termbox.GetCell(s.CX, s.CY).Ch))
}

func (s *State) MoveDown() {
	if len(s.Lines)-1 != s.CY {
		s.CY += 1
		if s.CX > len(s.Lines[s.CY])-1 {
			s.CX = 0
		}
	} else {
		return
	}

	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()
	s.Logger.Log(fmt.Sprintf("x: %d, y: %d | length line: %d | current char: %s rune: %d", s.CX, s.CY, len(s.Lines[s.CY]), string(termbox.GetCell(s.CX, s.CY).Ch), termbox.GetCell(s.CX, s.CY).Ch))
}
