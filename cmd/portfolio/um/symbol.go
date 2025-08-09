package um

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/um"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	symbolConfigCmd = &cobra.Command{
		Use:   "symbol",
		Short: "symbol config",
	}

	showSymbolConfigCmd = &cobra.Command{
		Use:   "show",
		Short: "Show symbol config",
		Long: `Get current account symbol configuration.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Futures-Symbol-Config`,
		Run: showSymbolConfig,
	}

	leverageSetCmd = &cobra.Command{
		Use:     "set-leverage",
		Aliases: []string{"leverage"},
		Short:   "set leverage",
		Long: `Change user's initial leverage of specific symbol in UM.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-UM-Initial-Leverage`,
		Run: setLeverage,
	}
)

func InitSymbolConfigCmds() []*cobra.Command {
	showSymbolConfigCmd.Flags().StringP("symbol", "s", "", "symbol")

	leverageSetCmd.Flags().StringP("symbol", "s", "", "symbol")
	leverageSetCmd.Flags().IntP("leverage", "l", 0, "leverage")
	leverageSetCmd.MarkFlagRequired("symbol")
	leverageSetCmd.MarkFlagRequired("leverage")

	symbolConfigCmd.AddCommand(showSymbolConfigCmd, leverageSetCmd)
	return []*cobra.Command{symbolConfigCmd}
}

func showSymbolConfig(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	resp, err := client.GetUMSymbolConfig(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&resp)
}

func setLeverage(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	resp, err := client.SetUMLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("leverage set success, symbol: %s, leverage: %d, maxNotionalValue: %s\n", resp.Symbol, resp.Leverage, resp.MaxNotionalValue)
}
