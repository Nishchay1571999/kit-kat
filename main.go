package main

import (
	"github.com/Nishchay1571999/kit-kat/input"
	"github.com/Nishchay1571999/kit-kat/terminal"
	"github.com/Nishchay1571999/kit-kat/utils"
)

func main() {
	utils.EditorClearScreen()
	terminal.EnableRawMode()
	defer terminal.RestoreState()

	input.ReadInputLoop()
}
