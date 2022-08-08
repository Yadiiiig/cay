package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/yadiiiig/cay/core"
	"github.com/yadiiiig/cay/file"
	"github.com/yadiiiig/cay/logger"
)

func init() {
	err := termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	var state core.State

	logger, err := logger.Create("logs.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	state.Logger = logger
	state.Logger.Clean()

	err = file.Read("test.txt", &state)
	if err != nil {
		fmt.Println(err)
		return
	}

	//go misc.Blink(&state)

	for {
		ev := termbox.PollEvent()
		if done := state.InputCapture(&ev); done {
			break
		}

		termbox.SetCursor(state.CX, state.CY)
		termbox.Flush()
	}

	termbox.Close()
}
