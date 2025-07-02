package unified

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/spf13/cobra"
)

var (
	commissionRateCmd = &cobra.Command{
		Use:   "um-commission-rate",
		Short: "unified UM commission rate",
		Run:   commissionRate,
	}
)

func init() {
	commissionRateCmd.Flags().StringP("symbol", "s", "", "symbol")
	commissionRateCmd.MarkFlagRequired("symbol")
}

func InitCommissionRateCmds() []*cobra.Command {
	return []*cobra.Command{commissionRateCmd}
}

func commissionRate(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")

	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	commissionRate, err := client.GetCommissionRate(symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s commission mraker rate: %s, taker rate: %s\n", symbol, commissionRate.MakerCommissionRate, commissionRate.TakerCommissionRate)
}
