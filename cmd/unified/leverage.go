package unified

import (
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

func InitLeverageCmds() []*cobra.Command {
	leverageCmd.AddCommand(leverageSetCmd)
	return []*cobra.Command{leverageCmd}
}

func setLeverage(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	leverage, _ := cmd.Flags().GetInt("leverage")
	err := client.SetUMLeverage(symbol, leverage)
	if err != nil {
		log.Fatal(err)
	}
}
