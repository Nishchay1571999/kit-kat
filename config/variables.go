package config

import "github.com/Nishchay1571999/kit-kat/types"

var (
	TerminalHeight = 0
	TerminalWidth  = 0
	Cursor         = types.CursorPos{
		X: 0,
		Y: 0,
	}
)

func SetTerminalHeight(height int, width int) {
	TerminalHeight = height
	TerminalWidth = width
}

func SetCursorPosition(currPos types.CursorPos) {
	Cursor = currPos
}
