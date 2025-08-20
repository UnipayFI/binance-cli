package futures

import (
	"fmt"
	"log"
	"os"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	binancefutures "github.com/adshao/go-binance/v2/futures"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Support create, cancel, list futures orders",
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Get all account orders; active, canceled, or filled.",
		Long: `Get all account orders; active, canceled, or filled.
- These orders will not be found
		- order status is 'CANCELED' or 'EXPIRED' AND order has NO filled trade AND created time + 3 days < current time
		- order create time + 90 days < current time

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/All-Orders`,
		Run: orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "Get all open orders on a symbol",
		Long: `Get all open orders on a symbol.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Current-All-Open-Orders`,
		Run: orderOpenList,
	}
	orderForceCloseCmd = &cobra.Command{
		Use:   "force",
		Short: "Query user's Force Orders",
		Long: `Query user's Force Orders.
- If "autoCloseType" is not sent, orders with both of the types will be returned
- If "startTime" is not sent, data within 7 days before "endTime" can be queried

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Users-Force-Orders`,
		Run: forceCloseOrder,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a new order",
		Long: `Create a new order.
* support all docs parameters

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api`,
		Run: createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "cancel order",
		Long: `cancel order 
If either orderId or orgClientOrderId is provided, the specified order will be canceled. 
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-Order
Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Cancel-All-Open-Orders`,
		Run: cancelOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")

	orderListCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderListCmd.MarkFlagRequired("symbol")

	orderForceCloseCmd.Flags().StringP("symbol", "s", "", "symbol")
	orderForceCloseCmd.Flags().StringP("autoCloseType", "t", "", "\"LIQUIDATION\" for liquidation orders, \"ADL\" for ADL orders.")
	orderForceCloseCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderForceCloseCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderForceCloseCmd.Flags().Int64P("limit", "l", 50, "limit, max 100")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "side")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "type")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}
	orderCreateCmd.MarkFlagRequired("symbol")

	orderCancelCmd.Flags().Int64P("orderId", "i", 0, "orderId")
	orderCancelCmd.Flags().StringP("origClientOrderId", "c", "", "origClientOrderId")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderForceCloseCmd, orderCreateCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	orders, err := client.GetOrderList(symbol, limit, startTime, endTime, orderID)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func forceCloseOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	autoCloseType, _ := cmd.Flags().GetString("autoCloseType")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	orders, err := client.GetForceOrders(symbol, binancefutures.ForceOrderCloseType(autoCloseType), startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderID)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	clientOrderID, _ := cmd.Flags().GetString("origClientOrderId")

	var err error
	if orderID == 0 && clientOrderID == "" {
		err = client.CancelAllOrders(symbol)
	} else {
		err = client.CancelOrder(symbol, orderID, clientOrderID)
	}
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order canceled")
	}
}
