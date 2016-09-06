package main

import (
	"log"
	"os"

	_ "github.com/gerardmrk/feeder/matchers"
	"github.com/gerardmrk/feeder/search"
)

// init is called prior to main.
func init() {
	log.SetOutput(os.Stdout)
}

// main is the entrypoint for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("ISIS")
}
