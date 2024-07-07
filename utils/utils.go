// utils/utils.go
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // ANSI escape sequence to clear screen
	}
}

// Die prints an error message to stderr and exits the program with status 1
func Die(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(1)
}
