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
	positionSideCmd = &cobra.Command{
		Use:   "position-side",
		Short: "futures position side",
	}

	positionSideStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "show position side status",
		Run:     positionSideStatus,
	}

	positionSideStatusChangeCmd = &cobra.Command{
		Use:     "change",
		Aliases: []string{"c"},
		Short:   "change position side status",
		Run:     positionSideStatusChange,
	}
)

func InitPositionSideCmds() []*cobra.Command {
	positionSideStatusChangeCmd.Flags().BoolP("dualSidePosition", "d", true, "change dual side position")
	positionSideStatusChangeCmd.MarkFlagRequired("dualSidePosition")

	positionSideCmd.AddCommand(positionSideStatusCmd, positionSideStatusChangeCmd)
	return []*cobra.Command{positionSideCmd}
}

func positionSideStatus(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	position, err := client.GetPositionSide()
	if err != nil {
		log.Fatalf("futures position side status error: %v", err)
	}
	fmt.Printf("dual side position: %v\n", position)
}

func positionSideStatusChange(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	dualSidePosition, _ := cmd.Flags().GetBool("dualSidePosition")
	err := client.ChangePositionSide(dualSidePosition)
	if err != nil {
		log.Fatalf("futures position side status change error: %v", err)
	}
	fmt.Printf("dual side position changed to: %v\n", dualSidePosition)
}
