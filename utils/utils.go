// utils/utils.go
package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Nishchay1571999/kit-kat/config"
	"golang.org/x/crypto/ssh/terminal"
)

// EditorClearScreen clears the terminal screen
func EditorClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // ANSI escape sequence to clear screen
	}
	editorRefreshScreen()
}
func EditorAddRows() {
	// Get the file descriptor for stdout
	fd := int(os.Stdout.Fd())

	// Retrieve the terminal size
	_, height, err := terminal.GetSize(fd)
	fmt.Println(err)
	config.SetTerminalHeight(height)
	if err != nil {
		Die("Could not find your Terminal")
	}
	// Define the string to write
	str := "~\r\n"

	// Convert the string to a byte slice
	bytes := []byte(str)

	for i := 0; i < height; i++ {
		fmt.Println(string(bytes))
	}

}
func editorRefreshScreen() {
	// ANSI escape sequence to clear the screen: "\x1b[2J"
	EditorClearScreen := "\x1b[2J"
	// ANSI escape sequence to move cursor to top-left: "\x1b[H"
	moveCursor := "\x1b[H"

	// Write the escape sequences to stdout (os.Stdout)
	if _, err := os.Stdout.Write([]byte(EditorClearScreen + moveCursor)); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to stdout: %v\n", err)
		os.Exit(1)
	}
}

// Die prints an error message to stderr and exits the program with status 1
func Die(message string) {
	log.Fatalf(message)
}
