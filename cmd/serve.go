package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"signable-sdk/http"
	"signable-sdk/model"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the SDK API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Building up data fixtures...")
		fmt.Print("\n")

		// setup the data fixtures in memory
		buildFixtures()

		// start http server
		fmt.Println("Starting SDK Server...")
		http.Init()
		fmt.Print("\n")

		return nil
	},
}

func buildFixtures() {
	fmt.Println("Generating API Keys...")

	keys, err := model.InitKey()
	if err != nil {
		panic("failed creating API keys")
	}

	// output the keys
	fmt.Printf("⇨ Success Key: %v \n", keys[0].Token)
	fmt.Printf("⇨ Error Key: %v \n", keys[1].Token)
	fmt.Println("\n")
}
