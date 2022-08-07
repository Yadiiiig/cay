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
	lines := make(map[int]string)

	for snr.Scan() {
		txt := snr.Text()

		lines[y] = txt

		for i := 0; i < len(txt); i++ {
			termbox.SetChar(i, y, rune(txt[i]))
		}

		y++
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

	// if len(lines) != 0 {
	// 	state.LL = len(lines[0])
	// } else {
	// 	state.LL = 0
	// }

	termbox.SetCursor(s.CX, s.CY)
	termbox.Flush()

	return nil
}
