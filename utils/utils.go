package utils

import (
	"fmt"
	"os"
)

// Die prints an error message to stderr and exits the program with status 1
func Die(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(1)
}
