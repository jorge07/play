package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long: "Display current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1")
		},
	})
}
