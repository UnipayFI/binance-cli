package futures

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	binancefutures "github.com/adshao/go-binance/v2/futures"
	"github.com/spf13/cobra"
)

var (
	symbolConfigCmd = &cobra.Command{
		Use:   "symbol-config",
		Short: "Show symbol config and set margin type",
	}

	showSymbolConfigCmd = &cobra.Command{
		Use:   "show",
		Short: "Show symbol config",
		Long: `Get current account symbol configuration.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Symbol-Config`,
		Run: showSymbolConfig,
	}

	setSymbolConfigCmd = &cobra.Command{
		Use:   "set",
		Short: "Set margin type",
		Long: `Change symbol level margin type.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Margin-Type`,
		Run: setMarginType,
	}
)

func InitSymbolConfigCmds() []*cobra.Command {
	showSymbolConfigCmd.Flags().StringP("symbol", "s", "", "symbol")

	setSymbolConfigCmd.Flags().StringP("symbol", "s", "", "symbol")
	setSymbolConfigCmd.Flags().StringP("marginType", "m", "", "margin type, ISOLATED or CROSSED")
	setSymbolConfigCmd.MarkFlagRequired("symbol")
	setSymbolConfigCmd.MarkFlagRequired("marginType")

	symbolConfigCmd.AddCommand(showSymbolConfigCmd, setSymbolConfigCmd)
	return []*cobra.Command{symbolConfigCmd}
}

func showSymbolConfig(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	symbolConfig, err := client.GetSymbolConfig(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&symbolConfig)
}

func setMarginType(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	marginType, _ := cmd.Flags().GetString("marginType")
	err := client.SetMarginType(symbol, binancefutures.MarginType(marginType))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("margin type set to", marginType)
}
