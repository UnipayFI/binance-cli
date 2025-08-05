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
		Use: "order",
	}

	orderHistoryCmd = &cobra.Command{
		Use:   "history",
		Short: "list order history",
		Run:   orderHistory,
	}
	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "create order",
		Run:     createOrder,
	}
	orderCancelCmd = &cobra.Command{
		Use:     "rm",
		Aliases: []string{"cancel"},
		Short:   "cancel order",
		PreRun: func(cmd *cobra.Command, _ []string) {
			orderID, _ := cmd.Flags().GetString("orderID")
			clientOrderID, _ := cmd.Flags().GetString("clientOrderID")
			if orderID == "" && clientOrderID == "" {
				log.Fatal("orderID or clientOrderID is required")
			}
		},
		Run: cancelOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("symbol", "s", "", "symbol")
	orderCmd.MarkFlagRequired("symbol")

	orderHistoryCmd.Flags().Int64P("orderID", "i", 0, "orderID")
	orderHistoryCmd.Flags().IntP("limit", "l", 500, "limit, max 1000")
	orderHistoryCmd.Flags().Int64P("startTime", "a", 0, "start time")
	orderHistoryCmd.Flags().Int64P("endTime", "e", 0, "end time")

	var side, orderType string
	orderCreateCmd.Flags().StringVarP(&side, "side", "S", "", "side")
	orderCreateCmd.Flags().StringVarP(&orderType, "type", "t", "", "type")
	orderCreateCmd.FParseErrWhitelist = cobra.FParseErrWhitelist{
		UnknownFlags: true,
	}

	orderCancelCmd.Flags().StringP("orderID", "i", "", "orderID")
	orderCancelCmd.Flags().StringP("clientOrderID", "c", "", "clientOrderID")

	orderCmd.AddCommand(orderHistoryCmd, orderCreateCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func orderHistory(cmd *cobra.Command, args []string) {
	client := spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	limit, _ := cmd.Flags().GetInt("limit")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	orderID, _ := cmd.Flags().GetInt64("orderID")
	orders, err := client.GetOrderHistory(symbol, orderID, startTime, endTime, limit)
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
	orderID, _ := cmd.Flags().GetInt64("orderID")
	clientOrderID, _ := cmd.Flags().GetString("clientOrderID")
	err := client.CancelOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order canceled")
	}
}
