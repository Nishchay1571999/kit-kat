// input/input.go
package input

import (
	"fmt"
	"os"

	"github.com/Nishchay1571999/kit-kat/terminal"
	"github.com/Nishchay1571999/kit-kat/utils"
)

// ReadInputLoop handles reading input and processing it
func ReadInputLoop() {
	utils.EditorAddRows()
	var buffer = make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buffer)
		HandleBuffer(buffer, n)
		if err != nil || n == 0 {
			utils.EditorClearScreen()
			fmt.Println(err)
			utils.Die("Error while taking Input")
		}
	}
}

/*
TODO:
1. Handle the Next Line Problem
2. Change the way Print is working

*/

func HandleBuffer(buffer []byte, n int) {
	terminal.MoveCursorTo()
	if terminal.EditorShouldQuit(buffer) {
		utils.EditorClearScreen()
		utils.Die("Quitting Kit Kat")
	} else {
		switch string(buffer[:n]) {
		case "\n":
			terminal.NextLine()
		default:
			fmt.Print(string(buffer[:n]))

		}
	}
}
