package unified

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	incomeCmd = &cobra.Command{
		Use:   "um-income",
		Short: "um-income",
		Run:   umIncome,
	}
)

func InitIncomeCmds() []*cobra.Command {
	return []*cobra.Command{incomeCmd}
}

func umIncome(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, err := client.GetUMIncome()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
