package margin

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
		Short: "Interest history",
		Long: `Get Margin Borrow/Loan Interest History.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Margin-BorrowLoan-Interest-History`,
		Run: interestHistory,
	}
)

func InitInterestHistoryCmds() []*cobra.Command {
	interestHistoryCmd.Flags().StringP("asset", "a", "", "asset")
	interestHistoryCmd.Flags().Int64P("startTime", "t", 0, "start time")
	interestHistoryCmd.Flags().Int64P("endTime", "e", 0, "end time")
	interestHistoryCmd.Flags().Int64P("current", "c", 1, "Currently querying page. Start from 1")
	interestHistoryCmd.Flags().Int64P("size", "s", 10, "page size, max 100")
	interestHistoryCmd.Flags().BoolP("archived", "r", false, "Set to true for archived data from 6 months ago")

	return []*cobra.Command{interestHistoryCmd}
}

func interestHistory(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	asset, _ := cmd.Flags().GetString("asset")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	current, _ := cmd.Flags().GetInt64("current")
	size, _ := cmd.Flags().GetInt64("size")
	archived, _ := cmd.Flags().GetBool("archived")
	list, err := client.GetMarginInterestHistory(asset, startTime, endTime, current, size, archived)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
