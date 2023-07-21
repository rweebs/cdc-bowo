/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/rweebs/cdc-bowo/internal/app/config"
	"github.com/rweebs/cdc-bowo/internal/app/lib"
	"github.com/rweebs/cdc-bowo/internal/app/services"
	"github.com/spf13/cobra"
)

// startCmd represents the add command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start change-dll or blue-green",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var startChangeDDLCmd = &cobra.Command{
	Use:   "change-ddl",
	Short: "Change DDL",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Waiting Replication Catch Up Before stoppping replication")
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
		cdcMgmt.StartSync()
	},
}

var startBlueGreenCmd = &cobra.Command{
	Use:   "blue-green",
	Short: "Start Blue Green Deployment",
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
		cdcMgmt.StartBlueGreen()
	},
}

func init() {
	startCmd.AddCommand(startChangeDDLCmd)
	startCmd.AddCommand(startBlueGreenCmd)
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.PersistentFlags().StringP("config", "c", "./config.test.json", "supply the config path")
}
