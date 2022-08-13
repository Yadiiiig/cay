package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

type cell struct {
	cell_type rune
	width     int
	value     string
}

// a, b, c, x, y and z are indictor cells
// idea is that they are named the same in config
// a, b and c are left of screen. x, y and z are right
var cells []cell

// Initializes the status bar and draws the info for the first time
func (s *State) LoadBar() {

	width, height := termbox.Size()
	//
	//	status := "Hello status bar"
	//	for i := 0; i < len(status); i++ {
	//		termbox.SetChar(i, height-1, rune(status[i]))
	//	}

	// initialize the cell array with data from user
	cells = []cell{
		{
			cell_type: 'a',
			value:     "%FN%",
			width:     0,
		},
		{
			cell_type: 'b',
			value:     "B",
			width:     0,
		},
		{cell_type: 'c', value: "C", width: 0}, 
		{
			cell_type: 'x',
			value:     "[%CY%:%CX%]",
			width:     0,
		}, {cell_type: 'y', value: "", width: 0}, {cell_type: 'z', value: "", width: 0},
	}
	s.drawCell(width, height, &cells)

	// 	text := "[" + strconv.Itoa(s.CY) + ":" + strconv.Itoa(s.CX) + "]"
	// 	length := width - len(text)
	// 	for i := width - 1; i > length-1; i-- {
	// 		termbox.SetChar(i, height-1, rune(text[i-length]))
	//     }
	//     bar.xIndicator.width = length
}

// Updates the data on the bar
func (s *State) UpdateBar() {

	width, height := termbox.Size()

	//	status := "Status Hello"
	//	for i := 0; i < len(status); i++ {
	//		termbox.SetChar(i, height-1, rune(status[i]))
	//	}

	//  text := "[" + strconv.Itoa(s.CY) + ":" + strconv.Itoa(s.CX) + "]"
	// 	length := width - len(text)
	//     if length > bar.xIndicator.width {
	//         for i := width-1; i > bar.xIndicator.width-1; i-- {
	//             termbox.SetChar(i, height-1, rune(0))
	//         }
	//     }
	// 	for i := width - 1; i > length-1; i-- {
	// 		termbox.SetChar(i, height-1, rune(text[i-length]))
	//     }
	//     bar.xIndicator.width = length
	s.drawCell(width, height, &cells)
}

func (s *State) drawCell(width int, height int, cells *[]cell) {
	for i, cell := range *cells {
		//	regex := regexp.MustCompile("%[A-z]{2}%")
		//	text := cell.value
		//	for _, v := range regex.FindAllString(cell.value, -1) {
		//		value := s.calcItem(v)
		//		text = strings.Replace(text, v, value, 1)
		//	}
		text := s.calcItem(cell.value)
		if cell.cell_type == 'a' || cell.cell_type == 'b' || cell.cell_type == 'c' {
			// left side of bar
			length := len(text)
			offset := s.calc_offset((*cells)[:i])
			if len(text) < cell.width {
				for i := offset; i < cell.width-1+offset; i++ {
					termbox.SetChar(i, height-1, rune(0))
				}
			}
			for i := offset; i < offset+length; i++ {
				termbox.SetChar(i, height-1, rune(text[i-offset]))
			}
			(*cells)[i].width = len(text)
		} else {
			// right side of bar
			length := width - len(text)
			offset := s.calc_offset((*cells)[i+1:])
			if len(text) < cell.width {
				for i := width - 2 - offset; i > width-cell.width-1-offset; i-- {
					termbox.SetChar(i, height-1, rune('h'))
				}
			}
			for i := width - 1 - offset; i > length-1-offset; i-- {
				termbox.SetChar(i, height-1, rune(text[i-length+offset]))
			}
			(*cells)[i].width = len(text)
		}
	}
}

func (s *State) calcItem(item string) string {
	regex := regexp.MustCompile("%[A-z]{2}%")
	for _, v := range regex.FindAllString(item, -1) {
		// value := s.calcItem(v)
		// text = strings.Replace(text, v, value, 1)
		switch v {
		case "%CX%":
			item = strings.Replace(item, v, strconv.Itoa(s.CX), 1)
		case "%CY%":
			item = strings.Replace(item, v, strconv.Itoa(s.CY), 1)
		case "%FN%":
			item = strings.Replace(item, v, s.Name, 1)
		}
	}
	return item
}

func (s *State) calc_offset(cells []cell) int {
	offset := 0
	for _, v := range cells {
		offset += len(s.calcItem(v.value))
	}
	return offset
}
