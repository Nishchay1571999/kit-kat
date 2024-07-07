package main

import (
	"github.com/Nishchay1571999/kit-kat/input"
	"github.com/Nishchay1571999/kit-kat/terminal"
	"github.com/Nishchay1571999/kit-kat/utils"
)

func main() {
	utils.ClearScreen()
	terminal.EnableRawMode()
	defer terminal.RestoreState()

	input.ReadInputLoop()
}
