package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balances",
	Run:   balances,
	Short: "show account balances",
}

func InitBalanceCmds() []*cobra.Command {
	return []*cobra.Command{balanceCmd}
}

func balances(cmd *cobra.Command, args []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	balances, err := client.GetBalances()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&balances)
}
