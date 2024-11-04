package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version string
)

var rootCmd = &cobra.Command{
	Use:   "wtf-gateway",
	Short: "Rest API to expose wtf backend to UI",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
