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
			s.Lines[s.CY] += string(key)

			s.CX += 1

			termbox.SetChar(s.CX, s.CY, value)
			s.Lines[s.CY] += string(key)
		} else {
			termbox.SetChar(s.CX, s.CY, key)
			s.Logger.Log(s.Lines[s.CY])
			s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + string(key) + s.Lines[s.CY][s.CX:]
			s.Logger.Log(s.Lines[s.CY])

			s.CX += 1

			termbox.SetChar(s.CX, s.CY, value)
			s.Lines[s.CY] = s.Lines[s.CY][:s.CX] + string(key) + s.Lines[s.CY][s.CX:]
			s.Logger.Log(s.Lines[s.CY])

			s.LoadLine()
		}
		//fmt.Fprintf(tvw, "%s%s", string(key), value)
	} else {
		termbox.SetChar(s.CX, s.CY, key)
		s.Lines[s.CY] += string(key)
		s.CX += 1
	}
}

// func KeyStroke(tvw *tview.TextView, key rune) {
// 	switch key {
// 	case 34:
// 		fmt.Fprintf(tvw, "%s%s", string(key), `"`)
// 	case 39:
// 		fmt.Fprintf(tvw, "%s%s", string(key), "'")
// 	case 40:
// 		fmt.Fprintf(tvw, "%s%s", string(key), ")")
// 		// Set highlight position in the middle of those 2 brackets
// 	case 91:
// 		fmt.Fprintf(tvw, "%s%s", string(key), "]")
// 	case 123:
// 		fmt.Fprintf(tvw, "%s%s", string(key), "}")
// 	default:
// 		fmt.Fprintf(tvw, "%s%d", string(key), key)
// 	}
// }
