package misc

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/yadiiiig/cay/core"
)

func Blink(state *core.State) {
	d := 250 * time.Millisecond
	i := 0
	for range time.Tick(d) {
		if i == 1000 {
			i = 0
		}
		if i%2 == 0 {
			termbox.HideCursor()
		} else {
			termbox.SetCursor(state.CX, state.CY)
		}

		i++
	}
}
