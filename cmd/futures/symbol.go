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
		Use:   "symbol",
		Short: "symbol config(leverage & margin type)",
	}

	showSymbolConfigCmd = &cobra.Command{
		Use:   "show",
		Short: "Show symbol config",
		Long: `Get current account symbol configuration.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Symbol-Config`,
		Run: showSymbolConfig,
	}

	leverageSetCmd = &cobra.Command{
		Use:     "set-leverage",
		Aliases: []string{"leverage"},
		Short:   "set leverage",
		Run:     setLeverage,
	}

	setSymbolConfigCmd = &cobra.Command{
		Use:     "set-margin-type",
		Aliases: []string{"margin-type"},
		Short:   "Set margin type",
		Long: `Change symbol level margin type.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Margin-Type`,
		Run: setMarginType,
	}
)

func InitSymbolConfigCmds() []*cobra.Command {
	showSymbolConfigCmd.Flags().StringP("symbol", "s", "", "symbol")

	leverageSetCmd.Flags().StringP("symbol", "s", "", "symbol")
	leverageSetCmd.Flags().IntP("leverage", "l", 0, "leverage")
	leverageSetCmd.MarkFlagRequired("symbol")
	leverageSetCmd.MarkFlagRequired("leverage")

	setSymbolConfigCmd.Flags().StringP("symbol", "s", "", "symbol")
	setSymbolConfigCmd.Flags().StringP("marginType", "m", "", "margin type, ISOLATED or CROSSED")
	setSymbolConfigCmd.MarkFlagRequired("symbol")
	setSymbolConfigCmd.MarkFlagRequired("marginType")

	symbolConfigCmd.AddCommand(showSymbolConfigCmd, leverageSetCmd, setSymbolConfigCmd)
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

func setLeverage(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	err := client.SetLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("leverage set to %d for %s\n", leverage, symbol)
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
