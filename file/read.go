package file

import (
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/yadiiiig/cay/core"
)

// Read is designed to open a file and read it's content,
// together with various fields from core.State
func Read(file string, s *core.State) error {
	tmp, err := os.Open(file)
	if err != nil {
		return err
	}

	defer tmp.Close()

	snr := bufio.NewScanner(tmp)
	snr.Split(bufio.ScanLines)

	y := 0
	lines := [][]rune{{}}

	for snr.Scan() {
		txt := snr.Text()

		lines = append(lines, []rune(txt))
		// lines[y] = []rune(txt)

		for i := 0; i < len(lines[y]); i++ {
			termbox.SetChar(i, y-1, lines[y][i])
		}

		y++
	}
	for i := 0; i < len(lines[y]); i++ {
		termbox.SetChar(i, y-1, lines[y][i])
	}

	if len(lines) == 0 {
		lines = append(lines, []rune{})
	}

	stat, err := tmp.Stat()
	if err != nil {
		return err
	}

	s.Name = stat.Name()
	s.Path = file
	s.Size = int(stat.Size())

	s.CX = 0
	s.CY = 0

	s.Lines = lines

	_, height := termbox.Size()
	s.BY = height - 2

	s.LoadBar()

	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()

	return nil
}
