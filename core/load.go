package core

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

func (s *State) LoadLine() {
	w, _ := termbox.Size()

	for i := 0; i < w; i++ {
		termbox.SetChar(i, s.CY, 0)
	}

	for i := 0; i < len(s.Lines[s.CY]); i++ {
		termbox.SetChar(i, s.CY, rune(s.Lines[s.CY][i]))
	}
}

func (s *State) LoadIndexLine(length, index, y int) {
	s.Logger.Log(fmt.Sprintf("len %d, index %d, y %d", length, index, y))
	for i := index; i <= length; i++ {
		termbox.SetChar(i, y, 0)
	}
}

func (s *State) WriteIndexLine(length, index, y int) {
	s.Logger.Log(fmt.Sprintf("len %d, index %d, y %d", length, index, y))
	for i := index; i <= length-1; i++ {
		termbox.SetChar(i, y, s.Lines[y][i])
	}
}

func (s *State) LoadIndexRestLine(length, y int) {
	for i := y; i < len(s.Lines); i++ {
		if i >= s.BY+1 {
			s.UpdateBar()
			continue
		}
		for j := 0; j < len(s.Lines[i]); j++ {
			termbox.SetChar(j, i, s.Lines[i][j])
		}
	}
}

func (s *State) LoadIndexRestNewLine(length, index, y int) {
    w, _ := termbox.Size()
    for i := index; i <= length; i++ {
        termbox.SetChar(i, y-1, 0)
    }

    for i := y; i < len(s.Lines); i++ {
        for j := 0; j < w; j++ {
            termbox.SetChar(j, i, 0)
        }
    }

    for i := y; i < len(s.Lines); i++ {
        if i >= s.BY+1 {
            s.UpdateBar()
            continue
        }
        for j := 0; j < len(s.Lines[i]); j++ {
            termbox.SetChar(j, i, s.Lines[i][j])
        }
    }
}

func (s *State) RemoveLines(length, y int) {
	w, _ := termbox.Size()

	for i := y+1; i < length; i++ {
		for j := 0; j < w; j++ {
			termbox.SetChar(j, i, 0)
		}
	}
} 