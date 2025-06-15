package futures

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/spf13/cobra"
)

var (
	leverageCmd = &cobra.Command{
		Use:   "leverage",
		Short: "futures leverage",
	}
	leverageSetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"s"},
		Short:   "set leverage",
		Run:     setLeverage,
	}
)

func InitLeverageCmds() []*cobra.Command {
	leverageCmd.AddCommand(leverageSetCmd)
	leverageSetCmd.Flags().StringP("symbol", "s", "", "symbol")
	leverageSetCmd.Flags().IntP("leverage", "l", 0, "leverage")
	leverageSetCmd.MarkFlagRequired("symbol")
	leverageSetCmd.MarkFlagRequired("leverage")
	return []*cobra.Command{leverageCmd}
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
