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
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use: "order",
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list orders",
		Run:     listOrders,
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

	orderLeverageCmd = &cobra.Command{
		Use:     "leverage",
		Aliases: []string{"lv"},
		Short:   "leverage order",
		Run:     leverageOrder,
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

	orderLeverageCmd.Flags().IntP("leverage", "l", 1, "leverage")
	orderLeverageCmd.MarkFlagRequired("leverage")

	orderCmd.AddCommand(orderListCmd, orderCreateCmd, orderCancelCmd, orderLeverageCmd)
	return []*cobra.Command{orderCmd}
}

func listOrders(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	orders, err := client.GetOrders(symbol)
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
	orderID, _ := cmd.Flags().GetInt64("orderID")
	clientOrderID, _ := cmd.Flags().GetString("clientOrderID")
	err := client.CancelOrder(symbol, orderID, clientOrderID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("order canceled")
	}
}

func leverageOrder(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	symbolLeverage, err := client.LeverageOrder(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("leverage order:", symbolLeverage)
	}
}
