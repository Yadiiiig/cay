package core

import "github.com/yadiiiig/cay/logger"

type State struct {
	Name     string
	Path     string
	Size     int
	LastSave int

	// Cursor position
	CX, CY int

	Lines map[int]string

	Logger logger.Logger
}
