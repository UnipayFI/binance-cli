package spot

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/spot"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var dividendCmd = &cobra.Command{
	Use:   "dividend",
	Short: "Get dividend information",
}

var dividendListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Query asset dividend record",
	Long: `Query asset dividend record.

Docs Link: https://developers.binance.com/docs/wallet/asset/assets-divided-record`,
	Run: listDividends,
}

func InitDividendCmds() []*cobra.Command {
	dividendListCmd.Flags().StringP("asset", "a", "", "asset")
	dividendListCmd.Flags().Int64P("startTime", "s", 0, "start time")
	dividendListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	dividendListCmd.Flags().IntP("limit", "l", 20, "limit, max 500")

	dividendCmd.AddCommand(dividendListCmd)
	return []*cobra.Command{dividendCmd}
}

func listDividends(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")

	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	dividends, err := client.GetDividendHistory(asset, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&dividends)
}
