package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

type cell struct {
    width int
    value string
}

// a, b, c, x, y and z are indictor cells
// idea is that they are named the same in config
// a, b and c are left of screen. x, y and z are right
var bar struct {
    aIndicator cell
    bIndicator cell
    cIndicator cell
    xIndicator cell
    yIndicator cell
    zIndicator cell
}

// Initializes the status bar and draws the info for the first time
func (s *State) LoadBar() {

    width, height := termbox.Size()
//
//	status := "Hello status bar"
//	for i := 0; i < len(status); i++ {
//		termbox.SetChar(i, height-1, rune(status[i]))
//	}

    bar.xIndicator = cell{
        value: "[%CY%:%CX%]",
        width: 0,
    }
    s.drawCell(width, height, &bar.xIndicator)

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
     s.drawCell(width, height, &bar.xIndicator)
}

func (s *State) drawCell(width int, height int, cell *cell) {
    regex := regexp.MustCompile("%[A-z]{2}%")
    text := cell.value
    for _, v := range regex.FindAllString(cell.value, -1) {
        value := s.calcItem(v)
        text = strings.Replace(text, v, value, 1)
    }
	length := width - len(text)
    if length > cell.width {
        for i := width-1; i > cell.width-1; i-- {
            termbox.SetChar(i, height-1, rune(0))
        }
    }
	for i := width - 1; i > length-1; i-- {
		termbox.SetChar(i, height-1, rune(text[i-length]))
    }
    cell.width = length
}

func (s *State) calcItem(item string) string {
    switch item {
    case "%CX%":
        return strconv.Itoa(s.CX)
    case "%CY%":
        return strconv.Itoa(s.CY)
    default:
        return ""
    }
}
