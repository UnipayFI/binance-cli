package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use: "balance",
	Run: balance,
}

func InitBalanceCmds() []*cobra.Command {
	return []*cobra.Command{balanceCmd}
}

func balance(cmd *cobra.Command, args []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(balance)
}
