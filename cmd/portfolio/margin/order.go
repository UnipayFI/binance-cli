package margin

import (
	"fmt"
	"log"
	"os"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/margin"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Support create, cancel, list margin orders",
	}
	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Query All Margin Account Orders",
		Long: `Query All Margin Account Orders.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Margin-Account-Orders`,
		Run: orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "Query Current Margin Open Order",
		Long: `Query Current Margin Open Order.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-Margin-Open-Order`,
		Run: orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create margin order",
		Long: `New Margin Order.
* support all docs parameters

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-Margin-Order`,
		Run: createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel Margin Account Order",
		Long: `Cancel Margin Account Order.
If either orderId or orgClientOrderId is provided, the specified order will be canceled.
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-Margin-Account-Order
Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-Margin-Account-All-Open-Orders-on-a-Symbol`,
		Run: cancelOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderListCmd.Flags().StringP("symbol", "s", "", "symbol")
	orderListCmd.Flags().Int64P("orderID", "i", 0, "orderID")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.MarkFlagRequired("symbol")

	orderOpenListCmd.Flags().StringP("symbol", "s", "", "symbol")
	orderOpenListCmd.MarkFlagRequired("symbol")

	orderCreateCmd.Flags().StringP("symbol", "s", "", "symbol")
	orderCreateCmd.Flags().StringP("side", "S", "", "side")
	orderCreateCmd.Flags().StringP("type", "t", "", "type")
	orderCreateCmd.MarkFlagRequired("symbol")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}

	orderCancelCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderCancelCmd.Flags().StringP("orderID", "i", "", "orderID")
	orderCancelCmd.Flags().StringP("clientOrderID", "c", "", "clientOrderID")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderCreateCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	orders, err := client.GetOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	orders, err := client.GetOpenOrderList(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderID)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	clientOrderID, _ := cmd.Flags().GetString("clientOrderID")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	err := client.CancelOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("order canceled")
}
