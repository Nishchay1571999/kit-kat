package config

import "github.com/Nishchay1571999/kit-kat/types"

var (
	TerminalHeight = 0
	Cursor         = types.CursorPos{
		X: 0,
		Y: 0,
	}
)

func SetTerminalHeight(height int) {
	TerminalHeight = height
}

func SetCursorPosition(currPos types.CursorPos) {
	Cursor.X = currPos.X
	Cursor.X = currPos.Y
}
