package spot

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/spot"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "show account info",
		Long: `Get current account information.

Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/testnet/rest-api/account-endpoints#account-information-user_data`,
		Run: showAccount,
	}
)

func InitAccountCmds() []*cobra.Command {
	return []*cobra.Command{accountCmd}
}

func showAccount(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	account, err := client.GetAccountInfo()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&account)
}
