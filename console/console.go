package console

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/techievee/deputy/indexer"

	"github.com/briandowns/spinner"

	"github.com/manifoldco/promptui"
	"github.com/techievee/deputy/utilities"
)

// Console Main user object that references the index processor to process the user query
type Console struct {
	indexer *indexer.Indexer

	welcomePrompt     *promptui.Prompt
	operationSelector *promptui.Select
	searchPrompt      *promptui.Prompt
	continuePrompt    *promptui.Prompt
}

// NewConsole Initialize the new console objects
func NewConsole(indexer *indexer.Indexer) *Console {

	// Template for the entryPrompt to display
	entryPrompt := promptui.Prompt{
		Label:       "Welcome to Deputy Subordinates finder tool",
		Default:     " Press any key to continue...",
		Validate:    nil,
		HideEntered: true,
	}

	// Template for the user to select the option
	operationSelector := promptui.Select{
		Label: "Select search term:",
		Items: []string{"Find Subordinates", "Quit"},
	}

	// Template for the input fields for the user to input user id to query
	searchPrompt := promptui.Prompt{
		Label:       "Enter the User ID to search",
		Validate:    nil,
		HideEntered: true,
	}

	// Template for the hold the screen after query completion
	continuePrompt := promptui.Prompt{
		Label:       "Records found",
		Default:     " Press any key to continue",
		Validate:    nil,
		HideEntered: true,
	}

	console := &Console{
		indexer,
		&entryPrompt,
		&operationSelector,
		&searchPrompt,
		&continuePrompt,
	}
	return console
}

// RunInteractive Provides interactive console for the user
func (c *Console) RunInteractive() {

	// clear the initial texts
	utilities.CallClear()

	// runs welcome prompt
	c.welcomeText()
	utilities.CallClear()
	// Exit if the Exit is selected in the first screen
	for {
		// Select the operation
		opindex, _, err := c.operationPrompt()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if opindex == 1 {
			return
		}

		// Provide prompt for entering the user id
		userId, err := c.Prompt()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		// Find and print the subordinates
		if err := c.PrintSubordinates(userId); err != nil {
			fmt.Printf("Printing Subordinates failed %v\n", err)
		}

		// Prompt screen, to hold the results
		label := ""
		_, _ = c.continueTextPrompt(label)

		// clear for new search
		utilities.CallClear()

	}

}

// operationPrompt Prompt for operation prompt for the interactive console
func (c *Console) operationPrompt() (int, string, error) {
	return c.operationSelector.Run()
}

// welcomeText Runs the welcome prompt
func (c *Console) welcomeText() {

	_, _ = c.welcomePrompt.Run()
	return
}

// Prompt Prompt for search term input  selector based on previous selection
func (c *Console) Prompt() (string, error) {

	c.searchPrompt.Label = fmt.Sprintf("Enter the User ID to find all subordinates")
	c.searchPrompt.Validate = c.validate
	retStr, err := c.searchPrompt.Run()
	return retStr, err

}

// validate Function to validate the live data while user keying in user id
func (c *Console) validate(input string) error {

	if strings.TrimSpace(input) == "" {
		return errors.New("User ID is required")
	}

	if ValidateInputString(input) == false {
		return errors.New("Only number is allowed")
	}
	return nil

}

// ValidateInputString Validate that input is number and not zero
func ValidateInputString(input string) bool {

	userId, err := strconv.ParseUint(strings.TrimSpace(input), 10, 64)
	if err != nil || userId == 0 {
		// Invalid data, omit the row
		return false
	}

	return true
}

// continueTextPrompt Prompt for holding the search texts
func (c *Console) continueTextPrompt(labelText string) (string, error) {
	c.continuePrompt.Label = labelText
	return c.continuePrompt.Run()
}

// RunNonInteractive Runs without any input from the user based on the command line arguments
func (c *Console) RunNonInteractive(userId string) {
	utilities.CallClear()

	// Validate the input
	if err := c.validate(userId); err != nil {
		fmt.Println(err)
		return
	}

	if err := c.PrintSubordinates(userId); err != nil {
		fmt.Println(err)
	}
	label := fmt.Sprintf("\nPress any key to continue")
	_, _ = c.continueTextPrompt(label)
}

// PrintSubordinates PrintSubordinates for records from the index
func (c *Console) PrintSubordinates(userId string) error {

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	_ = s.Color("green")
	s.Prefix = "Calculating subordinates, please wait... "
	s.Start()
	defer s.Stop()

	var intUserId uint64
	intUserId, err := strconv.ParseUint(strings.TrimSpace(userId), 10, 64)
	if err != nil {
		return err
	}

	// PrintSubordinates for the search user
	utilities.CallClear()
	buffer := c.indexer.PrintSubordinates(intUserId)
	fmt.Println(string(buffer[:]))
	return nil
}
