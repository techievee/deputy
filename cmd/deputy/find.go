package deputy

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/techievee/deputy"
)

// Global variables used by the console
var userId string // Used for non-interactive console

// searchCmd represents the search sub cmd for non-interactive console
var searchCmd = &cobra.Command{
	Use:   "subordinates",
	Short: "Deputy CLI Application to Index and Search Subordinates",
	Long: `
Deputy subordinates is a CLI library that indexes and search JSON file.
This application is a tool that reads the json file, indexes it and 
searches for subordinates without any user interactions.`,
	Run: func(cmd *cobra.Command, args []string) {
		deputy.Main(roleFilePath, userFilePath, userId, false)
	},
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(searchCmd)
	// Here you will define your flags and configuration settings.
	// local flags which will only run when this cmd
	searchCmd.Flags().StringVarP(&userId, "user-id", "u", "", "ID of the user, for whom the subordinates need to be calculated")

	// Validate inputs
	if strings.Trim(userId, " ") != "" {
		cobra.CheckErr(fmt.Errorf("User Id not provided"))
	}

}
