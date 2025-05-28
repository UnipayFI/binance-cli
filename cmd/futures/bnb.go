package futures

import (
	"fmt"
	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/futures"
	"github.com/spf13/cobra"
	"log"
)

var (
	bnbCmd = &cobra.Command{
		Use: "bnb",
	}

	bnbStatusCmd = &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "show fee burn status",
		Run:     bnbBurnStatus,
	}

	bnbStatusChangeCmd = &cobra.Command{
		Use:     "change",
		Aliases: []string{"c"},
		Short:   "change fee burn status",
		Run:     bnbBurnStatusChange,
	}
)

func InitBnbCmds() []*cobra.Command {
	bnbStatusChangeCmd.Flags().BoolP("status", "s", true, "change status")
	bnbStatusChangeCmd.MarkFlagRequired("status")

	bnbCmd.AddCommand(bnbStatusCmd, bnbStatusChangeCmd)
	return []*cobra.Command{bnbCmd}
}

func bnbBurnStatus(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	status, err := client.BnbBurnStatus()
	if err != nil {
		log.Fatalf("futures fee burn status error: %v", err)
	}
	fmt.Printf("fee burn switch status: %v", status.FeeBurn)
}

func bnbBurnStatusChange(cmd *cobra.Command, _ []string) {
	client := futures.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	status, _ := cmd.Flags().GetBool("status")
	err := client.BnbBurnStatusChange(status)
	if err != nil {
		log.Fatalf("futures fee burn status change error: %v", err)
	}
	fmt.Println("fee burn status changed")
}
