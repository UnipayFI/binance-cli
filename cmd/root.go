package cmd

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "binance-cli",
	Short: "binance API for CLI version",
}

func init() {
	initCommandConfig()
	checkAPIKey()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func initCommandConfig() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func checkAPIKey() {
	if config.Config.APIKey == "" || config.Config.APISecret == "" {
		log.Fatal("API_KEY and API_SECRET must be set")
	}
}
