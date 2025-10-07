package cmd

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	universaltransfer "github.com/UnipayFI/binance-cli/exchange/universal_transfer"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	universalTransferCmd = &cobra.Command{
		Use:   "universal-transfer",
		Short: "Transfer asset and history",
	}

	universalTransferListCmd = &cobra.Command{
		Use:   "ls",
		Short: "list universal transfer history",
		Run:   universalTransferList,
		Long: `Query User Universal Transfer History.

Docs Link: https://developers.binance.com/docs/wallet/asset/query-user-universal-transfer`,
	}

	transferCmd = &cobra.Command{
		Use:   "transfer",
		Short: "transfer asset",
		PreRun: func(cmd *cobra.Command, args []string) {
			transferType, _ := cmd.Flags().GetString("type")
			asset, _ := cmd.Flags().GetString("asset")
			amount, _ := cmd.Flags().GetString("amount")
			if transferType == "" || asset == "" || amount == "" {
				log.Fatal("transferType, asset, amount are required")
			}
			if transferType == "ISOLATEDMARGIN_MARGIN" || transferType == "ISOLATEDMARGIN_ISOLATEDMARGIN" {
				from, _ := cmd.Flags().GetString("fromSymbol")
				to, _ := cmd.Flags().GetString("toSymbol")
				if from == "" || to == "" {
					log.Fatal("if transferType is ISOLATEDMARGIN_MARGIN or ISOLATEDMARGIN_ISOLATEDMARGIN, fromSymbol and toSymbol are required")
				}
			}
		},
		Long: `User universal transfer.
* You need to enable 'Permits Universal Transfer' option for the API Key which requests this endpoint.

Docs Link: https://developers.binance.com/docs/wallet/asset/user-universal-transfer`,
		Run: universalTransfer,
	}
)

func init() {
	universalTransferListCmd.Flags().String("type", "", "transfer type")
	universalTransferListCmd.Flags().Int64("startTime", 0, "start time")
	universalTransferListCmd.Flags().Int64("endTime", 0, "end time")
	universalTransferListCmd.Flags().Int64("current", 0, "current")
	universalTransferListCmd.Flags().Int64("size", 100, "size, max 100")
	universalTransferListCmd.Flags().String("fromSymbol", "", "from symbol")
	universalTransferListCmd.Flags().String("toSymbol", "", "to symbol")

	transferCmd.Flags().String("type", "", "transfer type")
	transferCmd.Flags().String("asset", "", "asset")
	transferCmd.Flags().String("amount", "", "amount")
	transferCmd.Flags().String("fromSymbol", "", "from symbol")
	transferCmd.Flags().String("toSymbol", "", "to symbol")
	universalTransferCmd.AddCommand(universalTransferListCmd, transferCmd)
	RootCmd.AddCommand(universalTransferCmd)
}

func universalTransferList(cmd *cobra.Command, args []string) {
	transferType, _ := cmd.Flags().GetString("type")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	current, _ := cmd.Flags().GetInt64("current")
	size, _ := cmd.Flags().GetInt64("size")
	fromSymbol, _ := cmd.Flags().GetString("fromSymbol")
	toSymbol, _ := cmd.Flags().GetString("toSymbol")

	client := universaltransfer.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, err := client.List(transferType, startTime, endTime, current, size, fromSymbol, toSymbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}

func universalTransfer(cmd *cobra.Command, args []string) {
	transferType, _ := cmd.Flags().GetString("type")
	asset, _ := cmd.Flags().GetString("asset")
	amount, _ := cmd.Flags().GetString("amount")
	fromSymbol, _ := cmd.Flags().GetString("fromSymbol")
	toSymbol, _ := cmd.Flags().GetString("toSymbol")

	client := universaltransfer.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	err := client.Transfer(transferType, asset, amount, fromSymbol, toSymbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transfer successful")
}
