package main

import (
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init invoked before main
func init() {
	// Set the log output to standard output
	log.SetOutput(os.Stdout)
}

// main is the entry point for the application
func main() {
	// Run the search
	search.Run("president") // Define the missing function
}
