package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	commissionRateCmd = &cobra.Command{
		Use:     "commission-rate",
		Aliases: []string{"cr"},
		Short:   "show commission rate",
		Long: `Get User Commission Rate.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/User-Commission-Rate`,
		Run: showCommissionRate,
	}
)

func InitCommissionRateCmds() []*cobra.Command {
	commissionRateCmd.Flags().StringP("symbol", "s", "", "symbol")
	commissionRateCmd.MarkFlagRequired("symbol")
	return []*cobra.Command{commissionRateCmd}
}

func showCommissionRate(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	commissionRate, err := client.GetCommissionRate(symbol)
	if err != nil {
		log.Fatalf("futures commission rate error: %v", err)
	}
	printer.PrintTable(&commissionRate)
}
