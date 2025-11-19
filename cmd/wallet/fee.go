package wallet

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/wallet"
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
		Short:   "Get BNB Burn Status",
		Long: `Get BNB Burn Status.

Docs Link: https://developers.binance.com/docs/margin_trading/account/Get-BNB-Burn-Status`,
		Run: feeBurnStatus,
	}

	feeBurnStatusSetCmd = &cobra.Command{
		Use:     "set",
		Aliases: []string{"c"},
		Short:   "Toggle BNB Burn On Spot Trade And Margin Interest",
		Long: `Toggle BNB Burn On Spot Trade And Margin Interest.

Docs Link: https://developers.binance.com/docs/wallet/asset/Toggle-BNB-Burn-On-Spot-Trade-And-Margin-Interest`,
		Run: setFeeBurnStatus,
	}
)

func InitFeeCmds() []*cobra.Command {
	feeBurnStatusSetCmd.Flags().BoolP("spotBNBBurn", "s", false, "Determines whether to use BNB to pay for trading fees on SPOT")
	feeBurnStatusSetCmd.Flags().BoolP("interestBNBBurn", "i", false, "Determines whether to use BNB to pay for margin loan's interest")
	feeBurnStatusSetCmd.MarkFlagRequired("status")

	feeCmd.AddCommand(feeBurnStatusCmd, feeBurnStatusSetCmd)
	return []*cobra.Command{feeCmd}
}

func feeBurnStatus(cmd *cobra.Command, _ []string) {
	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	burnStatus, err := client.GetFeeBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn spotBNBBurn: %v, interestBNBBurn: %v\n", burnStatus.SpotBNBBurn, burnStatus.InterestBNBBurn)
}

func setFeeBurnStatus(cmd *cobra.Command, _ []string) {
	spotBNBBurn, _ := cmd.Flags().GetBool("spotBNBBurn")
	interestBNBBurn, _ := cmd.Flags().GetBool("interestBNBBurn")
	client := wallet.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	burnStatus, err := client.SetFeeBurnStatus(spotBNBBurn, interestBNBBurn)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Printf("fee burn changed to spotBNBBurn: %v, interestBNBBurn: %v\n", burnStatus.SpotBNBBurn, burnStatus.InterestBNBBurn)
}
