/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rweebs/cdc-bowo/internal/app/config"
	"github.com/rweebs/cdc-bowo/internal/app/lib"
	"github.com/rweebs/cdc-bowo/internal/app/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var configuration config.Config
		var err error
		configPath, _ := cmd.Flags().GetString("config")
		if configPath == "" {
			configuration, err = config.LoadConfig("./config.test.json")
			if err != nil {
				panic(err)
			}
		}
		configuration, err = config.LoadConfig(configPath)
		if err != nil {
			panic(err)
		}
		cache := lib.NewCache(configuration.CacheConfig.Host, configuration.CacheConfig.Port, configuration.CacheConfig.Password)
		cdcMgmt := services.NewCDCMgmtService(cache)
		if args[0] == "change-ddl" {

			cdcMgmt.StartSync()
			fmt.Println("Waiting Replication Catch Up Before stoppping replication")
		} else {
			cdcMgmt.StartBlueGreen()
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.PersistentFlags().StringP("config", "c", "./config.test.json", "supply the config path")
}
