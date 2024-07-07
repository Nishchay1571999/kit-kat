package main

import (
	"fmt"
	"os"

	"github.com/Nishchay1571999/kit-kat/terminal"
	"github.com/Nishchay1571999/kit-kat/utils"
)

func main() {
	terminal.EnableRawMode()
	defer terminal.RestoreState()
	var buffer = make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buffer)
		if terminal.ChangeMode(buffer) {
			utils.Die("Quitting Kit Kat")
		}

		fmt.Print(string(buffer[:n]))
		if err != nil || n == 0 {
			utils.Die("Error while taking Input")
		}
	}

}
