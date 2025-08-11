package um

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/um"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	incomeCmd = &cobra.Command{
		Use:   "income",
		Short: "Get UM Income History",
		Long: `Get UM Income History.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Income-History`,
		Run: umIncome,
	}
)

func InitIncomeCmds() []*cobra.Command {
	incomeCmd.Flags().StringP("symbol", "s", "", "symbol")
	incomeCmd.Flags().StringP("incomeType", "t", "", "income type")
	incomeCmd.Flags().Int64P("startTime", "a", 0, "start time")
	incomeCmd.Flags().Int64P("endTime", "e", 0, "end time")
	incomeCmd.Flags().IntP("limit", "l", 100, "limit, max 1000")
	incomeCmd.Flags().IntP("page", "p", 0, "page")
	return []*cobra.Command{incomeCmd}
}

func umIncome(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	incomeType, _ := cmd.Flags().GetString("incomeType")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	page, _ := cmd.Flags().GetInt("page")
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	list, err := client.GetIncome(symbol, incomeType, startTime, endTime, limit, page)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
