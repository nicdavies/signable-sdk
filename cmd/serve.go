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
	// API Key fixtures
	fmt.Println("Generating API Keys...")

	keys, err := model.InitKey()
	if err != nil {
		panic("failed creating API keys")
	}

	fmt.Printf("⇨ Success Key: %v \n", keys[0].Token)
	fmt.Printf("⇨ Error Key: %v \n", keys[1].Token)
	fmt.Println("\n")

	// contacts fixtures
	fmt.Println("Generating Contacts...")

	contacts, err := model.InitContact()
	if err != nil {
		panic("failed creating contacts")
	}

	fmt.Printf("⇨ Contacts: %+v", contacts)
	fmt.Println("\n")

	// TODO branding fixtures
	// TODO: envelope fixtures
	// TODO: settings fixtures
	// TODO: team fixtures
	// TODO: template fixtures
	// TODO: user fixtures
	// TODO: webhook fixtures
}
