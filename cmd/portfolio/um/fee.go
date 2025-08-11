package um

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/um"
	"github.com/spf13/cobra"
)

var (
	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "BNB payment fee",
		Long:  `BNB payment fee.`,
	}

	feeBurnStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Get user's BNB Fee Discount for UM Futures",
		Long: `Get user's BNB Fee Discount for UM Futures (Fee Discount On or Fee Discount Off ).

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Get-UM-Futures-BNB-Burn-Status`,
		Run: feeBurnStatus,
	}

	feeBurnStatusSetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"c"},
		Short:   "Change user's BNB Fee Discount for UM Futures on EVERY symbol",
		Long: `Change user's BNB Fee Discount for UM Futures (Fee Discount On or Fee Discount Off ) on EVERY symbol.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Toggle-BNB-Burn-On-UM-Futures-Trade`,
		Run: setFeeBurnStatus,
	}
)

func InitFeeCmds() []*cobra.Command {
	feeBurnStatusSetCmd.Flags().BoolP("status", "s", true, "change status")
	feeBurnStatusSetCmd.MarkFlagRequired("status")

	feeCmd.AddCommand(feeBurnStatusCmd, feeBurnStatusSetCmd)
	return []*cobra.Command{feeCmd}
}

func feeBurnStatus(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, err := client.FeeBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn status: %v\n", status.FeeBurn)
}

func setFeeBurnStatus(cmd *cobra.Command, _ []string) {
	client := um.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, _ := cmd.Flags().GetBool("status")
	err := client.SetFeeBurnStatus(status)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Printf("fee burn changed to: %v\n", status)
}
