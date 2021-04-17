package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"signable-sdk/http"
	"signable-sdk/model"
	"time"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the SDK API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// start a spinner
		spinner, _ := pterm.DefaultSpinner.Start("Building up data fixtures")
		time.Sleep(time.Second * 1)

		// setup the data fixtures in memory
		_ = model.InitFixtures()

		// start http server
		spinner.UpdateText("Starting SDK Server")
		time.Sleep(time.Second * 1)
		spinner.Success()

		http.Init()

		return nil
	},
}
