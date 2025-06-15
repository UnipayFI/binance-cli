package unified

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/spf13/cobra"
)

var (
	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "show fee burn status and change it",
	}

	feeBurnStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "show fee burn status",
		Run:     feeBurnStatus,
	}

	feeBurnStatusChangeCmd = &cobra.Command{
		Use:     "change",
		Aliases: []string{"c"},
		Short:   "change fee burn status",
		Run:     feeBurnStatusChange,
	}
)

func InitFeeCmds() []*cobra.Command {
	feeBurnStatusChangeCmd.Flags().BoolP("status", "s", true, "change status")
	feeBurnStatusChangeCmd.MarkFlagRequired("status")

	feeCmd.AddCommand(feeBurnStatusCmd, feeBurnStatusChangeCmd)
	return []*cobra.Command{feeCmd}
}

func feeBurnStatus(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	status, err := client.FeeBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn status: %v\n", status.FeeBurn)
}

func feeBurnStatusChange(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	status, _ := cmd.Flags().GetBool("status")
	err := client.FeeBurnStatusChange(status)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Printf("fee burn changed to: %v\n", status)
}
