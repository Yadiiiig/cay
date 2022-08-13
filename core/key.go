package core

import (
	"fmt"
	"strings"

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
	if s.CY == 0 && s.CX == 0 {
		return
	} else if s.CX != 0 {
		s.CX -= 1
		delete_in_line(&s.Lines[s.CY], s.CX)
		s.LoadLine()
	} else {
		s.CY -= 1
		s.CX = len(s.Lines[s.CY])
		s.Lines[s.CY] = append(s.Lines[s.CY], s.Lines[s.CY+1]...)

		s.LoadLine()
		s.RemoveLines(len(s.Lines), s.CY)
		delete_line(&s.Lines, s.CY+1)
	}
}

func (s *State) Delete() {
	s.Logger.Log(fmt.Sprintf("Length lines: %d, current index: %d", len(s.Lines), s.CY))
	// if Y is on last line   && if X is on the last char
	if len(s.Lines)-1 == s.CY && len(s.Lines[s.CY])-1 == s.CX {
		return
	} else if len(s.Lines[s.CY]) > s.CX {
		delete_in_line(&s.Lines[s.CY], s.CX+1)
		s.LoadLine()
	} else {
		prev_current := len(s.Lines[s.CY])
		s.RemoveLines(len(s.Lines), s.CY)
		s.Lines[s.CY] = append(s.Lines[s.CY], s.Lines[s.CY+1]...)

		delete_line(&s.Lines, s.CY+1)
		s.WriteIndexLine(len(s.Lines[s.CY]), prev_current, s.CY)
		s.LoadIndexRestLine(len(s.Lines), s.CY+1)
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

		s.CY += 1
		s.CX = 0

		s.LoadIndexRestNewLine(prev, len(s.Lines[s.CY-1]), s.CY)
	}
}

// calulate indecies of spaces in string ex. [4, 7, 9]
// move s.CX forward/backwards on those values on key event

func (s *State) WordSkipBackwards() {
	// if only 1 line	 && Y cursor is out of bounds
	if len(s.Lines) == 0 && s.CY > len(s.Lines) {
		return
	}

	// if X cursor is out of bounds
	if s.CX > len(s.Lines[s.CY]) {
		return
	}

	// looks for next index of a space from X cursor
	space := strings.Index(string(s.Lines[s.CY][s.CX+1:]), " ")

	// if there is no space following the X cursor
	// move X cursor to end of line
	if space == -1 {
		s.CX = 0
		return
	}
	// cursor is moved to the index after the space
	s.CX -= space - 2
}

func (s *State) WordSkipForwards() {
	// if only 1 line	 && Y cursor is out of bounds
	if len(s.Lines) == 0 && s.CY > len(s.Lines) {
		return
	}

	// if X cursor is out of bounds
	if s.CX > len(s.Lines[s.CY]) {
		return
	}

	// looks for next index of a space from X cursor
	space := strings.Index(string(s.Lines[s.CY][s.CX+1:]), " ")

	// if there is no space following the X cursor
	// move X cursor to end of line
	if space == -1 {
		s.CX = len(s.Lines[s.CY]) - 1
		return
	}
	// cursor is moved to the index after the space
	s.CX += space + 2
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

func all_indices(str, substr string) (indices []int) {
	if len(substr) == 0 {
		return
	}
	offset := 0
	for {
		i := strings.Index(str[offset:], substr)
		if i == -1 {
			return
		}
		offset += i
		indices = append(indices, offset)
		offset += len(substr)
	}
}
