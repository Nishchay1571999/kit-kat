// terminal/terminal.go
package terminal

import (
	"os"

	"golang.org/x/term"
)

var originalState *term.State

// EnableRawMode sets the terminal to raw mode
func EnableRawMode() {
	fd := int(os.Stdin.Fd())
	state, err := term.GetState(fd)
	if err != nil {
		panic(err)
	}
	originalState = state

	if _, err := term.MakeRaw(fd); err != nil {
		panic(err)
	}
}

// RestoreState restores the terminal to its original state
func RestoreState() {
	if originalState != nil {
		term.Restore(int(os.Stdin.Fd()), originalState)
	}
}

func ChangeMode(buffer []byte) bool {
	return buffer[0] == 3
}
