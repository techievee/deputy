package deputy

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	mainPackage "github.com/techievee/deputy"
	"github.com/techievee/deputy/utilities"
)

// Configuration file parameter
var cfgFile string

// Parameters used for both commands
var roleFilePath, userFilePath string

const (
	envPrefix = "DEPUTY"
)

// rootCmd represents the base command when called without any subcommands ( eg: flexera)
var rootCmd = &cobra.Command{
	Use:   "deputy",
	Short: "Deputy CLI Application to find subordinates for the user from the json File",
	Long: `
Deputy is a CLI library that indexes and search json file for finding all the subordinates for the user.
This application is a tool that reads the csv files, indexes it and 
calculates the total licences required for the application.`,

	Run: func(cmd *cobra.Command, args []string) {
		mainPackage.Main(roleFilePath, userFilePath, "", true)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

const (
	roleFileName = "roles.json"
	userFileName = "users.json"
)

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/deputy.yaml)")
	rootCmd.PersistentFlags().StringVarP(&roleFilePath, "roles", "R", "", fmt.Sprintf("Path of the JSON file to load Roles, Defaults to currentPath/%s ", roleFileName))
	rootCmd.PersistentFlags().StringVarP(&userFilePath, "users", "U", "", fmt.Sprintf("Path of the JSON file to load Users, Defaults to currentPath/%s ", userFileName))
}

// initConfig reads in config file and ENV variables if sets.
func initConfig() {

	v := viper.New()

	if cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".config" (without extension).
		v.AddConfigPath(home)
		v.SetConfigName("deputy.yaml")
	}

	// Get the current path where to look for files, if no flag present
	currentPath, err := utilities.GetCurrentPath()
	cobra.CheckErr(err)

	// Load default paths
	if roleFilePath == "" {
		roleFilePath = filepath.Join(currentPath, roleFileName)
	}

	if userFilePath == "" {
		userFilePath = filepath.Join(currentPath, userFileName)
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv() // read in environment variables that match
	bindFlags(rootCmd, v)
	bindFlags(searchCmd, v)
	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", v.ConfigFileUsed())
	}

}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
