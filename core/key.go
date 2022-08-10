package core

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

var SymbolKeys = map[rune]rune{
	34:  34,
	39:  39,
	40:  41,
	91:  93,
	123: 125,
}

func (s *State) KeyStrokeMap(key rune) {
	if value, ok := SymbolKeys[key]; ok {
		if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
			termbox.SetChar(s.CX, s.CY, key)
			s.Lines[s.CY] = append(s.Lines[s.CY], key)

			s.CX += 1

			termbox.SetChar(s.CX, s.CY, value)
			s.Lines[s.CY] = append(s.Lines[s.CY], value)
		} else {
			termbox.SetChar(s.CX, s.CY, key)
			s.Logger.Log(string(s.Lines[s.CY]))
			insert_in_line(&s.Lines[s.CY], key, s.CX)

			s.CX += 1

			termbox.SetChar(s.CX, s.CY, value)
			insert_in_line(&s.Lines[s.CY], value, s.CX)

			s.LoadLine()
		}
	} else {
		if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
			termbox.SetChar(s.CX, s.CY, key)
			s.Lines[s.CY] = append(s.Lines[s.CY], key)
			s.CX += 1

		} else {
			termbox.SetChar(s.CX, s.CY, key)
			insert_in_line(&s.Lines[s.CY], key, s.CX)

			s.CX += 1

			s.LoadLine()
		}
	}
}

func (s *State) AddSpace() {
	if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
		termbox.SetChar(s.CX, s.CY, 32)
		s.Lines[s.CY] = append(s.Lines[s.CY], 32)
		s.CX += 1

	} else {
		termbox.SetChar(s.CX, s.CY, 32)
		insert_in_line(&s.Lines[s.CY], 32, s.CX)

		s.CX += 1

		s.LoadLine()
	}
}

func (s *State) AddTab() {
	if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
		termbox.SetChar(s.CX, s.CY, 9)
		s.Lines[s.CY] = append(s.Lines[s.CY], 9)
		s.CX += 1

	} else {
		termbox.SetChar(s.CX, s.CY, 9)
		insert_in_line(&s.Lines[s.CY], 9, s.CX)

		s.CX += 1

		s.LoadLine()
	}
}

func (s *State) BackSpace() {
	// cant backspace on index 0
	// bring next lines up
	if s.CY == 0 && s.CX == 0 {
		return
	} else if s.CX == 0 {
		s.CY -= 1
		s.Logger.Log(string(s.Lines[s.CY]))
		s.Lines[s.CY] = append(s.Lines[s.CY], s.Lines[s.CY+1]...)
		s.LoadLine()
		s.RemoveLines(len(s.Lines), s.CY)
		s.Logger.Log(string(s.Lines[s.CY]))

	} else {
		s.CX -= 1
		delete_in_line(&s.Lines[s.CY], s.CX)
		s.LoadLine()
	}
}

func (s *State) Delete() {
	s.Logger.Log(fmt.Sprintf("Length lines: %d, current index: %d", len(s.Lines), s.CY))
	// if Y is on last line   && if X is on the last char	  || if X is 'outside' the last char
	if len(s.Lines)-1 == s.CY && len(s.Lines[s.CY])-1 == s.CX || len(s.Lines[s.CY]) == s.CX {
		return
	} else if (len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX) {
		prev_current := len(s.Lines[s.CY])
		// prev_next := len(s.Lines[s.CY+1])
		s.RemoveLines(len(s.Lines), s.CY)
		s.Lines[s.CY] = append(s.Lines[s.CY], s.Lines[s.CY+1]...)
		delete_line(&s.Lines, s.CY+1)
		
		s.WriteIndexLine(len(s.Lines[s.CY]), prev_current, s.CY)
		s.LoadIndexRestLine(len(s.Lines), s.CY+1)

	} else {
		delete_in_line(&s.Lines[s.CY], s.CX+1)
		s.LoadLine()
	}
}

func (s *State) NewLine() {
	if len(s.Lines)-1 == s.CY {
		if len(s.Lines[s.CY]) == s.CX && s.BY != s.CY {
			s.Lines = append(s.Lines, []rune{})
			s.CY += 1
			s.CX = 0
		} else {
			if s.CY == s.BY {
				return
			}
			prev := len(s.Lines[s.CY])

			s.Lines = append(s.Lines, s.Lines[s.CY][s.CX:])
			s.Lines[s.CY] = s.Lines[s.CY][:s.CX]

			s.CX = 0
			s.CY += 1

			s.LoadIndexLine(prev, len(s.Lines[s.CY-1]), s.CY-1)
			s.LoadLine()
		}
	} else {
		if s.CY == s.BY {
			return
		}
		prev := len(s.Lines[s.CY])

		insert_line(&s.Lines, s.Lines[s.CY][s.CX:], s.CY+1)
		s.Lines[s.CY] = s.Lines[s.CY][:s.CX]

		s.CX = 0
		s.CY += 1

		s.LoadIndexRestNewLine(prev, len(s.Lines[s.CY-1]), s.CY)
	}
}

// specific line operations

func insert_in_line(line *[]rune, key rune, i int) {
	*line = append(*line, 0)
	copy((*line)[i+1:], (*line)[i:])
	(*line)[i] = key

}

func delete_in_line(line *[]rune, i int) {
	copy((*line)[i:], (*line)[i+1:])
	(*line)[len((*line))-1] = 0
	*line = (*line)[:len(*line)-1]
}

// lines operation
func insert_line(lines *[][]rune, line []rune, i int) {
	*lines = append(*lines, []rune{})
	copy((*lines)[i+1:], (*lines)[i:])
	(*lines)[i] = line
}

func delete_line(lines *[][]rune, i int) {
	copy((*lines)[i:], (*lines)[i+1:])
	(*lines)[len((*lines))-1] = []rune{}
	*lines = (*lines)[:len(*lines)-1]
}
