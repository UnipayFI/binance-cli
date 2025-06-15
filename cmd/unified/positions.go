package unified

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	positionsCmd = &cobra.Command{
		Use:   "um-positions",
		Short: "unified UM positions",
	}

	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list positions",
		Run:     listUMPositions,
	}
)

func InitPositionsCmds() []*cobra.Command {
	positionsCmd.AddCommand(positionListCmd)
	return []*cobra.Command{positionsCmd}
}

func listUMPositions(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	positions, err := client.GetUMPositions()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&positions)
}
