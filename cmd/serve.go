package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"signable-sdk/http"
	"signable-sdk/model"
	"time"
)

var LogRequests bool

func init() {
	rootCmd.AddCommand(serveCmd)

	// register local command flags
	serveCmd.Flags().BoolVar(&LogRequests, "log-requests", false, "output http requests to the CLI")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the SDK API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// output "starting" status
		spinner, _ := pterm.DefaultSpinner.Start("Starting SDK")
		time.Sleep(time.Second * 2)

		// setup the data fixtures in memory
		spinner.UpdateText("Setting up data fixtures")
		fixtures := model.InitFixtures()
		time.Sleep(time.Second * 2)

		// display success message
		spinner.Success("Done!")

		// display fixture data
		pterm.FgLightGreen.Println("Up And Running:")
		printTokenFixtures(fixtures.Keys)

		// check if we should print request logs to CLI
		http.LogRequests = LogRequests

		// start http server
		http.Init()

		return nil
	},
}

func printTokenFixtures(fixtures []model.Key) {
	for _, e := range fixtures {
		pterm.Printf("%v key: %v\n", e.Type, e.Token)
	}
}
