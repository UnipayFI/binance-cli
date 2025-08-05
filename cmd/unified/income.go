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
	incomeCmd.Flags().StringP("symbol", "s", "", "symbol")
	incomeCmd.Flags().StringP("incomeType", "t", "", "income type")
	incomeCmd.Flags().Int64P("startTime", "a", 0, "start time")
	incomeCmd.Flags().Int64P("endTime", "e", 0, "end time")
	incomeCmd.Flags().IntP("limit", "l", 100, "limit, max 1000")
	return []*cobra.Command{incomeCmd}
}

func umIncome(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("incomeType")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, err := client.GetUMIncome(symbol, incomeType, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
