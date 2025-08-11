package futures

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	positionCmd = &cobra.Command{
		Use:   "position",
		Short: "show positions & show position risk & set position margin & change position side",
	}

	positionListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list positions",
		Long: `Get current account's all positions.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Account-Information-V3`,
		Run: listPositions,
	}

	positionRiskCmd = &cobra.Command{
		Use:     "risk",
		Aliases: []string{"r"},
		Short:   "show position risk",
		Long: `Get current position information(only symbol that has position or open orders will be returned).

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Position-Information-V3`,
		Run: showPositionRisk,
	}

	positionMarginCmd = &cobra.Command{
		Use:   "set-margin",
		Short: "set position margin",
		PreRun: func(cmd *cobra.Command, args []string) {
			typ, _ := cmd.Flags().GetString("type")
			if typ != "ADD" && typ != "REDUCE" {
				log.Fatalf("type must be ADD or REDUCE")
			}
		},
		Long: `Modify Isolated Position Margin

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Modify-Isolated-Position-Margin`,
		Run: setPositionMargin,
	}

	positionSideStatusCmd = &cobra.Command{
		Use:     "side",
		Aliases: []string{"s"},
		Short:   "Get user's position mode on EVERY symbol",
		Long: `Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-Current-Position-Mode`,
		Run: positionSideStatus,
	}

	positionSideStatusChangeCmd = &cobra.Command{
		Use:     "set-side",
		Aliases: []string{"c"},
		Short:   "Change Position Mode(TRADE)",
		Long: `Change Position Mode(TRADE).

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Change-Position-Mode`,
		Run: positionSideStatusChange,
	}
)

func InitPositionsCmds() []*cobra.Command {
	positionCmd.AddCommand(positionListCmd)

	positionRiskCmd.Flags().StringP("symbol", "s", "", "symbol")

	positionMarginCmd.Flags().StringP("symbol", "s", "", "symbol")
	positionMarginCmd.Flags().StringP("positionSide", "p", "BOTH", "Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent with Hedge Mode.")
	positionMarginCmd.Flags().Float64P("amount", "a", 0, "amount")
	positionMarginCmd.Flags().StringP("type", "t", "ADD", "ADD or REDUCE")
	positionMarginCmd.MarkFlagRequired("symbol")
	positionMarginCmd.MarkFlagRequired("amount")
	positionMarginCmd.MarkFlagRequired("type")

	positionSideStatusChangeCmd.Flags().BoolP("dualSidePosition", "d", true, "change dual side position")
	positionSideStatusChangeCmd.MarkFlagRequired("dualSidePosition")

	positionCmd.AddCommand(positionMarginCmd, positionRiskCmd, positionListCmd, positionSideStatusCmd, positionSideStatusChangeCmd)
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

func showPositionRisk(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	positions, err := client.GetPositionRisk(symbol)
	if err != nil {
		log.Fatalf("futures position risk error: %v", err)
	}
	printer.PrintTable(&positions)
}

func setPositionMargin(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	symbol, _ := cmd.Flags().GetString("symbol")
	positionSide, _ := cmd.Flags().GetString("positionSide")
	amount, _ := cmd.Flags().GetFloat64("amount")
	typ, _ := cmd.Flags().GetString("type")
	var t int
	if typ == "ADD" {
		t = 1
	} else {
		t = 2
	}
	err := client.ModifyPositionMargin(symbol, positionSide, amount, t)
	if err != nil {
		log.Fatalf("futures position margin set error: %v", err)
	}
	fmt.Printf("%s %s position %s %.6f\n", symbol, positionSide, typ, amount)
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
