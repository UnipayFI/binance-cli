package unified

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "unified order",
	}
	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list orders",
		Run:     orderList,
	}
	orderOpenListCmd = &cobra.Command{
		Use:   "um-open-list",
		Short: "list open orders",
		Run:   orderOpenList,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "um-create",
		Aliases: []string{"c"},
		Short:   "create UM order",
		Run:     createUMOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:   "um-cancel",
		Short: "cancel order",
		Long:  "cancel order \nIf either orderId or orgClientOrderId is provided, the specified order will be canceled. \nIf only the symbol is passed, all open orders for that trading pair will be canceled.",
		Run:   cancelOrder,
	}
	downloadOrderCmd = &cobra.Command{
		Use:     "um-download",
		Aliases: []string{"d"},
		Short:   "download UM order",
		Run:     downloadUMOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "side")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "type")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}
	orderCreateCmd.MarkFlagRequired("symbol")

	orderCancelCmd.Flags().StringP("orderID", "i", "", "orderID")
	orderCancelCmd.Flags().StringP("clientOrderID", "c", "", "clientOrderID")
	orderCancelCmd.MarkFlagRequired("symbol")

	orderListCmd.Flags().Int64P("orderID", "i", 0, "orderID")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderListCmd.MarkFlagRequired("symbol")

	orderOpenListCmd.MarkFlagRequired("symbol")

	downloadOrderCmd.Flags().Int64P("startTime", "a", 0, "start time")
	downloadOrderCmd.Flags().Int64P("endTime", "e", 0, "end time")

	orderCmd.AddCommand(orderListCmd, orderOpenListCmd, orderCreateCmd, orderCancelCmd, downloadOrderCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	limit, _ := cmd.Flags().GetInt("limit")
	orders, err := client.GetUMOrderList(symbol, orderID, startTime, endTime, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func orderOpenList(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetUMOpenOrders(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&orders)
}

func createUMOrder(cmd *cobra.Command, _ []string) {
	_, args, _ := cmd.Root().Find(os.Args[1:])
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	order, err := client.CreateUMOrder(common.ParseArgs(args))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order created, orderID:", order.OrderID)
	}
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	clientOrderID, _ := cmd.Flags().GetString("clientOrderID")

	var err error
	if orderID == 0 && clientOrderID == "" {
		err = client.CancelUMAllOrders(symbol)
	} else {
		err = client.CancelOrder(symbol, orderID, clientOrderID)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("order canceled: %v\n", orderID)
}

func downloadUMOrder(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, err := client.GetUMDownloadOrderID(symbol, startTime, endTime)
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
