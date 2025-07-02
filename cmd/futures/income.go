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
	incomeCmd = &cobra.Command{
		Use:   "income",
		Short: "show income history",
		Run:   income,
	}
)

func init() {
	incomeCmd.Flags().StringP("symbol", "s", "", "symbol")
	incomeCmd.Flags().StringP("income-type", "t", "", "income type")
	incomeCmd.Flags().Int64P("start-time", "a", 0, "start time")
	incomeCmd.Flags().Int64P("end-time", "e", 0, "end time")
	incomeCmd.Flags().Int64P("limit", "l", 0, "limit")
}

func InitIncomeCmds() []*cobra.Command {
	return []*cobra.Command{
		incomeCmd,
	}
}

func income(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("income-type")
	startTime, _ := cmd.Flags().GetInt64("start-time")
	endTime, _ := cmd.Flags().GetInt64("end-time")
	limit, _ := cmd.Flags().GetInt64("limit")

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	income, err := client.GetIncome(symbol, incomeType, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&income)
}
