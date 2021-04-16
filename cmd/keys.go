package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(keysCmd)
}

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Generate new SDK keys",
	Long:  "Generate new SDK keys. Generates both success keys, and error keys.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Generating API Keys...")

		return nil
	},
}
