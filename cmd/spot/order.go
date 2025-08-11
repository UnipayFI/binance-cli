package spot

import (
	"fmt"
	"log"
	"os"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/spot"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Support create, cancel, list orders",
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list orders",
		Run:     orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "list open orders",
		Long: `List open orders.

Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/testnet/rest-api/account-endpoints#current-open-orders-user_data`,
		Run: orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create order",
		Long: `Create a new order.
* support all docs parameters

Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-trade`,
		Run: createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "cancel order",
		Long: `cancel order 
If either orderId or orgClientOrderId is provided, the specified order will be canceled. 
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#cancel-order-trade
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

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderCreateCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderId")
	orders, err := client.GetOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderID)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
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
