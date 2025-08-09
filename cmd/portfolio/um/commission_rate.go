package um

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	commissionRateCmd = &cobra.Command{
		Use:   "commission-rate",
		Short: "Get User Commission Rate for UM",
		Long: `Get User Commission Rate for UM.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-User-Commission-Rate-for-UM`,
		Run: commissionRate,
	}
)

func InitCommissionRateCmds() []*cobra.Command {
	commissionRateCmd.Flags().StringP("symbol", "s", "", "symbol")
	commissionRateCmd.MarkFlagRequired("symbol")
	return []*cobra.Command{commissionRateCmd}
}

func commissionRate(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")

	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	commissionRate, err := client.GetCommissionRate(symbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s commission mraker rate: %s, taker rate: %s\n", symbol, commissionRate.MakerCommissionRate, commissionRate.TakerCommissionRate)
}
