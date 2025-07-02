package unified

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/spf13/cobra"
)

var (
	leverageCmd = &cobra.Command{
		Use:   "um-leverage",
		Short: "unified UM leverage",
	}
	leverageSetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"s"},
		Short:   "set leverage",
		Run:     setLeverage,
	}
)

func init() {
	leverageSetCmd.Flags().StringP("symbol", "s", "", "symbol")
	leverageSetCmd.Flags().IntP("leverage", "l", 0, "leverage")
	leverageSetCmd.MarkFlagRequired("symbol")
	leverageSetCmd.MarkFlagRequired("leverage")
}

func InitLeverageCmds() []*cobra.Command {
	leverageCmd.AddCommand(leverageSetCmd)
	return []*cobra.Command{leverageCmd}
}

func setLeverage(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	resp, err := client.SetUMLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("leverage set success, symbol: %s, leverage: %d, maxNotionalValue: %s\n", resp.Symbol, resp.Leverage, resp.MaxNotionalValue)
}
