package cmd

import (
	"errors"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "binance-cli",
	Short:   "Binance API for CLI version",
	PreRunE: checkAPIKey,
}

func init() {
	initCommandConfig()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func initCommandConfig() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func checkAPIKey(cmd *cobra.Command, args []string) error {
	if config.Config.APIKey == "" || config.Config.APISecret == "" {
		return errors.New("API_KEY and API_SECRET must be set")
	}
	return nil
}
