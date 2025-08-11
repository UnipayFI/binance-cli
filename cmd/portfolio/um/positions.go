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
	positionsCmd = &cobra.Command{
		Use:   "position",
		Short: "show positions & show position risk & change position side",
	}

	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list positions",
		Long: `Get current UM position information.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Account-Detail-V2`,
		Run: listUMPositions,
	}

	positionRiskCmd = &cobra.Command{
		Use:     "risk",
		Aliases: []string{"r"},
		Short:   "show position risk",
		Long: `Get current UM position information.
- Please use with user data stream "ACCOUNT_UPDATE" to meet your timeliness and accuracy needs.
- for One-way Mode user, the response will only show the "BOTH" positions
- for Hedge Mode user, the response will show "LONG", and "SHORT" positions.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-UM-Position-Information`,
		Run: showPositionRisk,
	}

	positionSideShowCmd = &cobra.Command{
		Use:     "side",
		Aliases: []string{"s"},
		Short:   "show position side",
		Long: `Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol in UM.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-UM-Current-Position-Mode`,
		Run: showPositionSide,
	}

	positionSideSetCmd = &cobra.Command{
		Use:     "set-side",
		Aliases: []string{"s"},
		Short:   "set position side",
		Long: `Change user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol in UM.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-UM-Position-Mode`,
		Run: setPositionSide,
	}
)

func InitPositionsCmds() []*cobra.Command {
	positionRiskCmd.Flags().StringP("symbol", "s", "", "symbol")

	positionSideSetCmd.Flags().BoolP("dualSidePosition", "d", true, "change dual side position")
	positionSideSetCmd.MarkFlagRequired("dualSidePosition")

	positionsCmd.AddCommand(positionListCmd, positionRiskCmd, positionSideShowCmd, positionSideSetCmd)
	return []*cobra.Command{positionsCmd}
}

func listUMPositions(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	positions, err := client.GetUMPositions()
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&positions)
}

func showPositionRisk(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	positions, err := client.GetUMPositionRisk(symbol)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&positions)
}

func showPositionSide(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	position, err := client.GetPositionSide()
	if err != nil {
		log.Fatalf("futures position side status error: %v", err)
	}
	fmt.Printf("dual side position: %v\n", position)
}

func setPositionSide(cmd *cobra.Command, _ []string) {
	client := um.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	dualSidePosition, _ := cmd.Flags().GetBool("dualSidePosition")
	err := client.ChangePositionSide(dualSidePosition)
	if err != nil {
		log.Fatalf("futures position side status change error: %v", err)
	}
	fmt.Printf("dual side position changed to: %v\n", dualSidePosition)
}
