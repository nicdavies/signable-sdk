package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"signable-sdk/http"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the SDK API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Starting SDK Server...")

		http.Init()

		return nil
	},
}
