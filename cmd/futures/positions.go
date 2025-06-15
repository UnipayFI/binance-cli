package futures

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	positionCmd = &cobra.Command{
		Use:   "positions",
		Short: "futures position",
	}

	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list positions",
		Run:     listPositions,
	}
)

func InitPositionsCmds() []*cobra.Command {
	positionCmd.AddCommand(positionListCmd)
	return []*cobra.Command{positionCmd}
}

func listPositions(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	positions, err := client.GetPositions()
	if err != nil {
		log.Fatalf("futures position list error: %v", err)
	}
	printer.PrintTable(&positions)
}
