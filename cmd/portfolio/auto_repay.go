package portfolio

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	autoRepayCmd = &cobra.Command{
		Use:   "auto-repay",
		Short: "Auto-repay",
	}

	autoRepayStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Query Auto-repay-futures Status",
		Long: `Query Auto-repay-futures Status.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Get-Auto-repay-futures-Status`,
		Run: autoRepayStatus,
	}
	setAutoRepayCmd = &cobra.Command{
		Use:   "set",
		Short: "Change Auto-repay-futures Status",
		Long: `Change Auto-repay-futures Status.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Change-Auto-repay-futures-Status`,
		Run: setAutoRepay,
	}
)

func init() {
	autoRepayCmd.PersistentFlags().BoolP("autoRepay", "a", false, "autoRepay")
	autoRepayCmd.MarkFlagRequired("autoRepay")
}

func InitAutoRepayCmds() []*cobra.Command {
	autoRepayCmd.AddCommand(autoRepayStatusCmd, setAutoRepayCmd)
	return []*cobra.Command{autoRepayCmd}
}

func autoRepayStatus(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, err := client.GetAutoRepayStatus()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("auto repay status: %v\n", status)
}

func setAutoRepay(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	status, _ := cmd.Flags().GetBool("autoRepay")
	s, err := client.SetAutoRepayStatus(status)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("auto repay status set to: %v\n", s)
}
