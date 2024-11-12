package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wtf version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
