package wallet

import (
	"fmt"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/wallet"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	dustCmd = &cobra.Command{
		Use:   "dust",
		Short: "dust asset conversion and history",
	}

	showDustCmd = &cobra.Command{
		Use:   "show",
		Short: "Get assets that can be converted into BNB",
		Long: `Get assets that can be converted into BNB.
		
Docs Link: https://developers.binance.com/docs/wallet/asset/assets-can-convert-bnb`,
		Run: showDust,
	}

	convertDustCmd = &cobra.Command{
		Use:   "convert",
		Short: "Convert dust assets to BNB.",
		Long: `Convert dust assets to BNB.
		
Docs Link: https://developers.binance.com/docs/wallet/asset/dust-transfer`,
		Run: convertDust,
	}

	historyDustCmd = &cobra.Command{
		Use:   "history",
		Short: "Dust conversion history",
		Long: `Dust conversion history.
* Only return last 100 records
* Only return records after 2020/12/01

Docs Link: https://developers.binance.com/docs/wallet/asset/dust-log`,
		Run: historyDust,
	}
)

func InitDustCmds() []*cobra.Command {
	showDustCmd.Flags().StringP("accountType", "a", "SPOT", "SPOT or MARGIN")

	convertDustCmd.Flags().StringP("asset", "s", "", "The asset list being converted. For example: BTC,USDT")
	convertDustCmd.Flags().StringP("accountType", "a", "SPOT", "SPOT or MARGIN")
	convertDustCmd.MarkFlagRequired("asset")

	historyDustCmd.Flags().Int64P("startTime", "s", 0, "start time")
	historyDustCmd.Flags().Int64P("endTime", "e", 0, "end time")

	dustCmd.AddCommand(showDustCmd, convertDustCmd, historyDustCmd)
	return []*cobra.Command{dustCmd}
}

func showDust(cmd *cobra.Command, args []string) {
	accountType, _ := cmd.Flags().GetString("accountType")
	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, totalBTC, totalBNB, percentage, err := client.ShowDust(accountType)
	if err != nil {
		fmt.Println("Error getting dust list:", err)
		return
	}
	printer.PrintTable(&list)
	fmt.Printf("\nTotal BTC: %s, Total BNB: %s, Percentage: %s\n", totalBTC, totalBNB, percentage)
}

func convertDust(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	accountType, _ := cmd.Flags().GetString("accountType")

	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, totalServiceCharge, totalTransfered, err := client.ConvertDust(asset, accountType)
	if err != nil {
		fmt.Println("Error converting dust:", err)
		return
	}
	printer.PrintTable(&list)
	fmt.Printf("\nTotal Service Charge: %s, Total Transfered: %s\n", totalServiceCharge, totalTransfered)
}

func historyDust(cmd *cobra.Command, args []string) {
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")

	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	list, err := client.HistoryDust(startTime, endTime)
	if err != nil {
		fmt.Println("Error getting history dust:", err)
		return
	}
	printer.PrintTable(&list)
}
