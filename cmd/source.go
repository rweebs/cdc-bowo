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
	Use:   "source",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.LoadConfig("./")
		if err != nil {
			panic(err)
		}
		// fmt.Println(config)
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		db := lib.NewDatabase(config.SourceConfig.Host, config.SourceConfig.Username, config.SourceConfig.Password, config.SourceConfig.Name, config.SourceConfig.Port)
		cache := lib.NewCache()
		cdcSourceService := services.NewCDCSourceServices(db, db, cache, config)
		go func() {
			<-c
			cdcSourceService.StopService()
			fmt.Println("Interrupt signal received, exiting program.")
			os.Exit(1)
		}()
		cdcSourceService.ExecuteDDLChange()
		cdcSourceService.StartService()

	},
}

func init() {
	rootCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
