package portfolio

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	bnbTransferCmd = &cobra.Command{
		Use:   "bnb-transfer",
		Short: "bnb-transfer",
		Run:   bnbTransfer,
	}
)

func init() {
	bnbTransferCmd.Flags().StringP("amount", "a", "", "amount")
	bnbTransferCmd.Flags().StringP("transferSide", "s", "", "transferSide: TO_UM、FROM_UM")
	bnbTransferCmd.MarkFlagRequired("amount")
	bnbTransferCmd.MarkFlagRequired("transferSide")
}

func InitBnbTransferCmds() []*cobra.Command {
	return []*cobra.Command{bnbTransferCmd}
}

func bnbTransfer(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	amount, _ := cmd.Flags().GetString("amount")
	transferSide, _ := cmd.Flags().GetString("transferSide")
	resp, err := client.BnbTransfer(amount, transferSide)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("bnb tranID: %v\n", resp.TranID)
}
