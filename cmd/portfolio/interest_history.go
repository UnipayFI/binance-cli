package portfolio

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	interestHistoryCmd = &cobra.Command{
		Use:   "interest-history",
		Short: "interest history",
		Run:   interestHistory,
	}
)

func InitInterestHistoryCmds() []*cobra.Command {
	return []*cobra.Command{interestHistoryCmd}
}

func interestHistory(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	asset, _ := cmd.Flags().GetString("asset")
	list, err := client.GetInterestHistory(asset)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
