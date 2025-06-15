package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var balancesCmd = &cobra.Command{
	Use:   "balances",
	Run:   balances,
	Short: "show account balances",
}

func InitBalancesCmds() []*cobra.Command {
	return []*cobra.Command{balancesCmd}
}

func balances(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(balances)
}
