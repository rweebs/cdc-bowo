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
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "report",
	Long:  `Command to report the status of the replication after the database schema change.`,
	Run: func(cmd *cobra.Command, args []string) {
		var configuration config.Config
		var err error
		configPath, _ := cmd.Flags().GetString("config")
		if configPath == "" {
			configuration, err = config.LoadConfig("./config.test.json")
			if err != nil {
				log.Panic(err)
			}
		}
		configuration, err = config.LoadConfig(configPath)
		if err != nil {
			log.Panic(err)
		}
		cache := lib.NewCache(configuration.CacheConfig.Host, configuration.CacheConfig.Port, configuration.CacheConfig.Password)
		dbSource := lib.NewDatabase(configuration.SourceConfig.Host, configuration.SourceConfig.Username, configuration.SourceConfig.Password, configuration.SourceConfig.Name, configuration.SourceConfig.Port)
		dbDest := lib.NewDatabase(configuration.DestConfig.Host, configuration.DestConfig.Username, configuration.DestConfig.Password, configuration.DestConfig.Name, configuration.DestConfig.Port)

		cdcReport := services.NewCDCReportServices(dbSource, dbDest, cache, configuration)
		cdcReport.GenerateReport()
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	reportCmd.PersistentFlags().StringP("config", "c", "./config.test.json", "supply the config path")
}
