package core

import (
	"github.com/nsf/termbox-go"
)

// Should be changed to runes instead of string
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
			//s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + string(key) + s.Lines[s.CY][s.CX:]
			//s.Logger.Log(s.Lines[s.CY])

			s.CX += 1

			termbox.SetChar(s.CX, s.CY, value)
			insert_in_line(&s.Lines[s.CY], value, s.CX)
			// s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + string(key) + s.Lines[s.CY][s.CX:]
			// s.Logger.Log(s.Lines[s.CY])

			s.LoadLine()
		}
		//fmt.Fprintf(tvw, "%s%s", string(key), value)
	} else {
		if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
			termbox.SetChar(s.CX, s.CY, key)
			s.Lines[s.CY] = append(s.Lines[s.CY], key)
			s.CX += 1

		} else {
			termbox.SetChar(s.CX, s.CY, key)
			insert_in_line(&s.Lines[s.CY], key, s.CX)
			//s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + string(key) + s.Lines[s.CY][s.CX:]
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
		//s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + " " + s.Lines[s.CY][s.CX:]
		s.CX += 1

		s.LoadLine()
	}
}

func (s *State) AddTab() {
	if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
		termbox.SetChar(s.CX, s.CY, 9)
		s.Lines[s.CY] = append(s.Lines[s.CY], 9)
		s.CX += 4

	} else {
		termbox.SetChar(s.CX, s.CY, 9)
		insert_in_line(&s.Lines[s.CY], 9, s.CX)
		//s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + " " + s.Lines[s.CY][s.CX:]
		s.CX += 4

		s.LoadLine()
	}
}

func (s *State) BackSpace() {
	if s.CX == 0 {
		switch s.CY {
		case 0:
			return
		default:
			s.CY -= 1
			s.CX = len(s.Lines[s.CY])
			return
		}
	}

	if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
		// s.Logger.Log(s.Lines[s.CY])
		s.Lines[s.CY] = s.Lines[s.CY][:len(s.Lines[s.CY])-1]
		s.CX -= 1
		termbox.SetChar(s.CX, s.CY, 0)
		// s.Logger.Log(s.Lines[s.CY])
	} else {
		// s.Logger.Log(s.Lines[s.CY])
		delete_in_line(&s.Lines[s.CY], s.CX-1)
		//s.Lines[s.CY] = delete_char([]rune(s.Lines[s.CY]), s.CX-1)
		//termbox.SetChar(s.CX, s.CY, 0)
		s.CX -= 1
		s.LoadLine()
		//s.Logger.Log(s.Lines[s.CY])
	}

}

func (s *State) Delete() {
	if len(s.Lines[s.CY]) == s.CX || len(s.Lines[s.CY])+1 == s.CX {
		return
	} else {
		//s.Logger.Log(s.Lines[s.CY])
		delete_in_line(&s.Lines[s.CY], s.CX)
		//termbox.SetChar(s.CX, s.CY, 0)
		s.LoadLine()
		//s.Logger.Log(s.Lines[s.CY])
	}
}

func (s *State) NewLine() {
	if len(s.Lines)-1 == s.CY {
		if len(s.Lines[s.CY]) == s.CX {
			s.Lines = append(s.Lines, []rune{})
			s.CY += 1
			s.CX = 0
		} else {
			prev := len(s.Lines[s.CY])

			s.Lines = append(s.Lines, s.Lines[s.CY][s.CX:])
			s.Lines[s.CY] = s.Lines[s.CY][:s.CX]

			s.CX = 0
			s.CY += 1

			s.LoadIndexLine(prev, len(s.Lines[s.CY-1]), s.CY-1)
			s.LoadLine()
		}
	} else {
		prev := len(s.Lines[s.CY])

		insert_line(&s.Lines, s.Lines[s.CY][s.CX:], s.CY+1)
		s.Lines[s.CY] = s.Lines[s.CY][:s.CX]

		s.CX = 0
		s.CY += 1

		s.LoadIndexRestLine(prev, len(s.Lines[s.CY-1]), s.CY)
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
