package cmd

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	universaltransfer "github.com/UnipayFI/binance-cli/exchange/universal_transfer"
	"github.com/spf13/cobra"
)

var (
	universalTransferCmd = &cobra.Command{
		Use:   "universal-transfer",
		Short: "User universal transfer",
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
		Long: `user universal transfer.
* You need to enable 'Permits Universal Transfer' option for the API Key which requests this endpoint.

Docs Link: https://developers.binance.com/docs/wallet/asset/user-universal-transfer`,
		Run: universalTransfer,
	}
)

func init() {
	universalTransferCmd.Flags().String("type", "", "transfer type")
	universalTransferCmd.Flags().String("asset", "", "asset")
	universalTransferCmd.Flags().String("amount", "", "amount")
	universalTransferCmd.Flags().String("fromSymbol", "", "from symbol")
	universalTransferCmd.Flags().String("toSymbol", "", "to symbol")
	RootCmd.AddCommand(universalTransferCmd)
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
