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
	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "BNB payment fee",
	}

	feeBurnStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Get BNB Burn Status (USER_DATA)",
		Long: `Get BNB Burn Status (USER_DATA).

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Get-BNB-Burn-Status`,
		Run: feeBurnStatus,
	}

	feeBurnStatusSetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"c"},
		Short:   "Change user's BNB Fee Discount on EVERY symbol",
		Long: `Change user's BNB Fee Discount (Fee Discount On or Fee Discount Off ) on EVERY symbol.

Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/account/rest-api/Toggle-BNB-Burn-On-Futures-Trade`,
		Run: feeBurnStatusChange,
	}
)

func InitFeeCmds() []*cobra.Command {
	feeBurnStatusSetCmd.Flags().BoolP("feeBurn", "b", true, `"true": Fee Discount On; "false": Fee Discount Off`)
	feeBurnStatusSetCmd.MarkFlagRequired("feeBurn")

	feeCmd.AddCommand(feeBurnStatusCmd, feeBurnStatusSetCmd)
	return []*cobra.Command{feeCmd}
}

func feeBurnStatus(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	status, err := client.FeeBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn status: %v\n", status.FeeBurn)
}

func feeBurnStatusChange(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	burn, _ := cmd.Flags().GetBool("feeBurn")
	err := client.FeeBurnStatusChange(burn)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Printf("fee burn changed to: %v\n", burn)
}
