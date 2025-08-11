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
		Short: "Query income history",
		Long: `Query income history.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Income-History`,
		Run: income,
	}
)

func init() {
	incomeCmd.Flags().StringP("symbol", "s", "", "symbol")
	incomeCmd.Flags().StringP("incomeType", "t", "", "income type")
	incomeCmd.Flags().Int64P("startTime", "a", 0, "Timestamp in ms to get funding from INCLUSIVE.")
	incomeCmd.Flags().Int64P("endTime", "e", 0, "Timestamp in ms to get funding until INCLUSIVE.")
	incomeCmd.Flags().Int64P("page", "p", 0, "page")
	incomeCmd.Flags().Int64P("limit", "l", 100, "limit, max 1000")
}

func InitIncomeCmds() []*cobra.Command {
	return []*cobra.Command{
		incomeCmd,
	}
}

func income(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("incomeType")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	page, _ := cmd.Flags().GetInt64("page")
	limit, _ := cmd.Flags().GetInt64("limit")

	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	income, err := client.GetIncome(symbol, incomeType, startTime, endTime, page, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&income)
}
