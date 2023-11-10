package cmd

import (
	"github.com/spf13/cobra"

	"github.com/mokhlesurr031/event-management-svc/internal/config"
)

var (
	rootCmd = &cobra.Command{
		Use:   "event-app",
		Short: "Backend Server",
		Long:  "Event Management Service backend API server using Golang",
	}
)

func init() {
	initConfig()
}

func initConfig() {
	config.Init()
}

func Execute() error {
	return rootCmd.Execute()
}
