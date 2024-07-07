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
	var buffer = make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buffer)
		if terminal.ChangeMode(buffer) {
			utils.ClearScreen()
			utils.Die("Quitting Kit Kat")
		}

		fmt.Print(string(buffer[:n]))
		if err != nil || n == 0 {
			utils.ClearScreen()
			utils.Die("Error while taking Input")
		}
	}
}
