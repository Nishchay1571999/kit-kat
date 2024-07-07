// terminal/terminal.go
package terminal

import (
	"os"

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

// ChangeMode checks if the input buffer signals a mode change
func ChangeMode(buffer []byte) bool {
	// Implement your logic for mode change detection here
	return false
}
