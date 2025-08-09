package um

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "UM Futures Fee",
		Long:  `UM Futures Fee.`,
	}

	feeBurnStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Get user's BNB Fee Discount for UM Futures",
		Long: `Get user's BNB Fee Discount for UM Futures (Fee Discount On or Fee Discount Off ).

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Get-UM-Futures-BNB-Burn-Status`,
		Run: feeBurnStatus,
	}

	feeBurnStatusChangeCmd = &cobra.Command{
		Use:     "change",
		Aliases: []string{"c"},
		Short:   "Toggle BNB Burn On UM Futures Trade (TRADE)",
		Long: `Toggle BNB Burn On UM Futures Trade (TRADE).

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Toggle-BNB-Burn-On-UM-Futures-Trade`,
		Run: feeBurnStatusChange,
	}
)

func InitFeeCmds() []*cobra.Command {
	feeBurnStatusChangeCmd.Flags().BoolP("status", "s", true, "change status")
	feeBurnStatusChangeCmd.MarkFlagRequired("status")

	feeCmd.AddCommand(feeBurnStatusCmd, feeBurnStatusChangeCmd)
	return []*cobra.Command{feeCmd}
}

func feeBurnStatus(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, err := client.FeeBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn status: %v\n", status.FeeBurn)
}

func feeBurnStatusChange(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, _ := cmd.Flags().GetBool("status")
	err := client.FeeBurnStatusChange(status)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Printf("fee burn changed to: %v\n", status)
}
