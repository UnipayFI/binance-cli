package portfolio

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var balancesCmd = &cobra.Command{
	Use:   "balances",
	Short: "Query account balance",
	Long: `Query account balance.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account`,
	Run: balances,
}

func InitBalancesCmds() []*cobra.Command {
	return []*cobra.Command{balancesCmd}
}

func balances(cmd *cobra.Command, args []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&balances)
}
