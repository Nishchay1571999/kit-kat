// input/input.go
package input

import (
	"fmt"
	"os"

	"github.com/Nishchay1571999/kit-kat/config"
	"github.com/Nishchay1571999/kit-kat/terminal"
	"github.com/Nishchay1571999/kit-kat/utils"
)

// ReadInputLoop handles reading input and processing it
func ReadInputLoop() {
	utils.EditorAddRows()
	terminal.MoveCursorTo()
	var buffer = make([]byte, config.TerminalWidth)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil || n == 0 {
			utils.EditorClearScreen()
			fmt.Println(err)
			utils.Die("Error while taking Input")
		}
		HandleBuffer(buffer, n)
	}
}

/*
TODO:
1. Handle the Next Line Problem
2. Change the way Print is working

*/

func HandleBuffer(buffer []byte, n int) {
	for i := 0; i < n; i++ {
		switch buffer[i] {
		case '\n':
			terminal.NextLine()
		case '\r':
			// Handle \r in case of Windows-style line endings (\r\n)
			if i+1 < n && buffer[i+1] == '\n' {
				terminal.NextLine()
				i++ // Skip the next \n
			} else {
				terminal.NextLine()
			}
		default:
			fmt.Print(string(buffer[i]))
		}
	}

	if terminal.EditorShouldQuit(buffer) {
		utils.EditorClearScreen()
		utils.Die("Quitting Kit Kat")
	}
}
