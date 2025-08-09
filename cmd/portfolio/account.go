package portfolio

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "show account info",
		Run:   account,
	}
)

func InitAccountCmds() []*cobra.Command {
	return []*cobra.Command{accountCmd}
}

func account(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&account)
}
