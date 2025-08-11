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
	tradeCmd = &cobra.Command{
		Use:   "trade",
		Short: "Get trades for a specific account and symbol.",
	}

	tradesListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Get trades for a specific account and symbol.",
		Long: `If 'startTime' and 'endTime' are both not sent, then the last 7 days' data will be returned.
The time between 'startTime' and 'endTime' cannot be longer than 7 days.
The parameter 'fromId' cannot be sent with 'startTime' or 'endTime'.
Only support querying trade in the past 6 months.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Account-Trade-List`,
		Run: trades,
	}
)

func InitTradesCmds() []*cobra.Command {
	tradesListCmd.Flags().StringP("symbol", "s", "", "symbol")
	tradesListCmd.MarkFlagRequired("symbol")

	tradesListCmd.Flags().StringP("orderId", "i", "", "orderId")
	tradesListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	tradesListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	tradesListCmd.Flags().StringP("fromId", "f", "", "fromId")
	tradesListCmd.Flags().Int64P("limit", "l", 500, "limit, max 1000")

	tradeCmd.AddCommand(tradesListCmd)
	return []*cobra.Command{tradeCmd}
}

func trades(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderId, _ := cmd.Flags().GetInt64("orderId")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	fromId, _ := cmd.Flags().GetInt64("fromId")
	limit, _ := cmd.Flags().GetInt("limit")
	trades, err := client.GetTrades(symbol, orderId, startTime, endTime, fromId, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&trades)
}
