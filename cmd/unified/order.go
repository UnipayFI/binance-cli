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
	orderCreateCmd = &cobra.Command{
		Use:     "um-create",
		Aliases: []string{"c"},
		Short:   "create UM order",
		Run:     createUMOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:     "cancel",
		Aliases: []string{"c"},
		Short:   "cancel order",
		Run:     cancelOrder,
	}
	downloadOrderCmd = &cobra.Command{
		Use:     "download",
		Aliases: []string{"d"},
		Short:   "download order",
		Run:     downloadOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderCmd.MarkFlagRequired("symbol")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "side")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "type")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}

	orderCancelCmd.Flags().StringP("orderID", "i", "", "orderID")
	orderCancelCmd.Flags().StringP("clientOrderID", "c", "", "clientOrderID")

	orderListCmd.Flags().Int64P("orderID", "i", 0, "orderID")
	orderListCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderListCmd.Flags().Int64P("endTime", "e", 0, "end time")
	orderListCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")

	downloadOrderCmd.Flags().Int64P("startTime", "a", 0, "start time")
	downloadOrderCmd.Flags().Int64P("endTime", "e", 0, "end time")

	orderCmd.AddCommand(orderListCmd, orderCreateCmd, orderCancelCmd, downloadOrderCmd)
	return []*cobra.Command{orderCmd}
}

func orderList(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
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
	err := client.CancelOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("order canceled: %v\n", orderID)
}

func downloadOrder(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, err := client.GetDownloadOrderID(symbol, startTime, endTime)
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
