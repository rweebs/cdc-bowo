/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rweebs/cdc-bowo/internal/app/config"
	"github.com/rweebs/cdc-bowo/internal/app/lib"
	"github.com/rweebs/cdc-bowo/internal/app/services"
	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "run",
	Short: "run the application using config.json or supply the config file using --config or -c flag",
	Long:  `run the application using config.json or supply the config file using --config or -c flag`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		var configuration config.Config
		var err error
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

		// fmt.Println(config)
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		sourceDb := lib.NewDatabase(configuration.SourceConfig.Host, configuration.SourceConfig.Username, configuration.SourceConfig.Password, configuration.SourceConfig.Name, configuration.SourceConfig.Port)
		destDb := lib.NewDatabase(configuration.DestConfig.Host, configuration.DestConfig.Username, configuration.DestConfig.Password, configuration.DestConfig.Name, configuration.DestConfig.Port)
		cache := lib.NewCache(configuration.CacheConfig.Host, configuration.CacheConfig.Port, configuration.CacheConfig.Password)
		cdcSourceService := services.NewCDCSourceServices(sourceDb, destDb, cache, configuration)
		cdcDestService := services.NewCDCDestServices(sourceDb, destDb, cache, configuration)
		go func() {
			<-c
			cdcSourceService.StopService()
			cdcDestService.StopService()
			fmt.Println("Interrupt signal received, exiting program.")
			os.Exit(1)
		}()
		timestampStopReplication, err := cdcSourceService.StopReplication()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Start Data Sync")
		cdcSourceService.ExecuteDDLChange()
		cdcSourceService.StartService(timestampStopReplication)
		cdcSourceService.StopService()
		startDestTimestamp, _ := cdcSourceService.GetTimeStampCutOff()
		fmt.Println("Start Blue Green Deployment")
		cdcDestService.StartService(startDestTimestamp)
		// cdcSourceService.StartService(0)
	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	sourceCmd.PersistentFlags().StringP("config", "c", "./config.test.json", "supply the config path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
