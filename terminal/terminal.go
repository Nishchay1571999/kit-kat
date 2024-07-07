// terminal/terminal.go
package terminal

import (
	"fmt"
	"os"

	"github.com/Nishchay1571999/kit-kat/config"
	"github.com/Nishchay1571999/kit-kat/types"
	"github.com/Nishchay1571999/kit-kat/utils"
	"golang.org/x/term"
)

var originalState *term.State

// EnableRawMode sets the terminal to raw mode
func EnableRawMode() {
	fd := int(os.Stdin.Fd())
	state, err := term.GetState(fd)
	if err != nil {
		utils.Die("Failed to get terminal state")
	}
	originalState = state

	if _, err := term.MakeRaw(fd); err != nil {
		utils.Die("Failed to set raw mode")
	}
}

// RestoreState restores the terminal to its original state
func RestoreState() {
	if originalState != nil {
		term.Restore(int(os.Stdin.Fd()), originalState)
	}
}
func MoveCursorTo() {
	// ANSI escape sequence to move cursor
	// Move cursor to position (x, y)
	fmt.Printf("\033[%d;%dH", config.Cursor.Y+1, config.Cursor.X+1) // Adding 1 because ANSI escape codes are 1-based
}
func NextLine() {
	nextPos := types.CursorPos{
		X: 0,
		Y: config.Cursor.Y + 1,
	}
	config.SetCursorPosition(nextPos)
	MoveCursorTo()
}

// EditorShouldQuit checks if the input buffer signals a mode change
func EditorShouldQuit(buffer []byte) bool {
	// Implement your logic for mode change detection here
	return buffer[0] == 3
}
