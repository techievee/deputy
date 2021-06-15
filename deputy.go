package deputy

import (
	"fmt"
	"os"

	"github.com/techievee/deputy/indexer"

	"github.com/techievee/deputy/console"
	"github.com/techievee/deputy/data"
)

// Main Core application logic that is initiated by the command processors from cmd package
func Main(rolesFile, usersFile, userId string, interactive bool) {

	// Read the data from the current specified  file to the memory
	//      fileHandle is always nil, used for running test cases
	d, err := data.NewDataStore(rolesFile, usersFile)
	if err != nil {
		fmt.Printf("\nError while openeing the file\n Error: %v", err)
		os.Exit(1)
	}

	err = d.ReadFiles()
	if err != nil {
		fmt.Printf("\nError while decoding the file\n Error: %v", err)
		os.Exit(1)
	}

	indexes := indexer.NewIndexer(d)
	// Generate indexes out of the data that was in memory
	indexes.GenerateIndex()

	// New Console application
	c := console.NewConsole(indexes)
	if interactive == true {
		// Invokes the interactive console
		MainInteractive(c)
	} else {
		// Invokes the non-interactive console
		MainNonInteractive(c, userId)
	}

}

// MainInteractive Runs the program in the interactive mode, asks for user input
func MainInteractive(c *console.Console) {
	c.RunInteractive()
}

// MainNonInteractive Runs the program in non-interactive mode, outputs the result and programs dies
func MainNonInteractive(c *console.Console, userId string) {
	c.RunNonInteractive(userId)
}
