package um

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/um"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Support create, cancel, list um orders",
	}
	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List orders",
		Long: `Get all account UM orders; active, canceled, or filled.
- These orders will not be found:
		- order status is 'CANCELED' or 'EXPIRED', AND 
		- order has NO filled trade, AND 
		- order create time + 3 days < current time

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-All-Margin-Account-Orders`,
		Run: orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "open",
		Short: "List open orders",
		Long: `Query current UM open order.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Query-Current-UM-Open-Order`,
		Run: orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create UM order",
		Long: `Create a new order.
* support all docs parameters

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/New-Margin-Order`,
		Run: createUMOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel order",
		Long: `cancel order 
If either orderId or orgClientOrderId is provided, the specified order will be canceled. 
If only the symbol is passed, all open orders for that trading pair will be canceled.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-UM-Order
Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Cancel-All-UM-Open-Orders`,
		Run: cancelOrder,
	}
	downloadOrderCmd = &cobra.Command{
		Use:     "download",
		Aliases: []string{"d"},
		Short:   "Download order history",
		Run:     downloadUMOrder,
	}
)

func InitOrderCmds() []*cobra.Command {

	var side, orderType string
	orderCreateCmd.Flags().StringP("symbol", "s", "", "symbol")
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "side")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "type")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}
	orderCreateCmd.MarkFlagRequired("symbol")

	orderCancelCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderCancelCmd.Flags().StringP("orderID", "i", "", "orderID")
	orderCancelCmd.Flags().StringP("clientOrderID", "c", "", "clientOrderID")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderListCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderListCmd.Flags().Int64P("orderID", "i", 0, "orderID")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.MarkFlagRequired("symbol")

	orderOpenListCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderOpenListCmd.MarkFlagRequired("symbol")

	downloadOrderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	downloadOrderCmd.Flags().Int64P("startTime", "a", 0, "start time")
	downloadOrderCmd.Flags().Int64P("endTime", "e", 0, "end time")

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderCreateCmd, orderCancelCmd, downloadOrderCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	orders, err := client.GetOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createUMOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	order, err := client.CreateOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderID)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	clientOrderID, _ := cmd.Flags().GetString("clientOrderID")

	var err error
	if orderID == 0 && clientOrderID == "" {
		err = client.CancelAllOrders(symbol)
	} else {
		err = client.CancelOrder(symbol, orderID, clientOrderID)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("order canceled: %v\n", orderID)
}

func downloadUMOrder(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, err := client.GetDownloadOrderID(startTime, endTime)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("downloadID:", orderID)
	for {
		time.Sleep(3 * time.Second)
		resp, err := client.GetDownloadOrderLink(orderID)
		if err != nil {
			log.Fatal(err)
		}
		if resp.Status == "completed" {
			fmt.Println("download link:", resp.URL)
			break
		} else if resp.Status == "processing" {
			fmt.Println("waiting for processing...")
		} else {
			log.Fatal("unknown status:", resp.Status)
		}
	}
}
