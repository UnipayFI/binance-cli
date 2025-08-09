package portfolio

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange/portfolio"
	"github.com/spf13/cobra"
)

var (
	repayFuturesNegativeBalanceCmd = &cobra.Command{
		Use:   "repay-futures-negative-balance",
		Short: "repay futures negative balance",
		Run:   repayFuturesNegativeBalance,
	}
)

func InitRepayFuturesNegativeBalanceCmds() []*cobra.Command {
	return []*cobra.Command{repayFuturesNegativeBalanceCmd}
}

func repayFuturesNegativeBalance(cmd *cobra.Command, _ []string) {
	client := portfolio.NewClient(config.Config.APIKey, config.Config.APISecret)
	resp, err := client.RepayFuturesNegativeBalance()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repay futures negative balance: %v\n", resp.Msg)
}
