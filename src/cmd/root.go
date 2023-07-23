/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
//create long description about the app

var rootCmd = &cobra.Command{
	Use:   "data-schema-change",
	Short: "Database schema change for blue green database deployment",
	Long: `Data Schema Change is a CLI application that will help you do
database schema change in Blue-Green Deployment in PostgreSQL.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdc-bowo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.AddCommand(addCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
