package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/futures"
	"github.com/spf13/cobra"
)

var futuresCmd = &cobra.Command{
	Use:   "futures",
	Short: "Futures",
}

func init() {
	futuresCmd.AddCommand(futures.InitAccountCmds()...)
	futuresCmd.AddCommand(futures.InitCommissionRateCmds()...)
	futuresCmd.AddCommand(futures.InitMultiAssetsModeCmds()...)
	futuresCmd.AddCommand(futures.InitOrderCmds()...)
	futuresCmd.AddCommand(futures.InitFeeCmds()...)
	futuresCmd.AddCommand(futures.InitIncomeCmds()...)
	futuresCmd.AddCommand(futures.InitPositionsCmds()...)
	futuresCmd.AddCommand(futures.InitSymbolConfigCmds()...)
	futuresCmd.AddCommand(futures.InitTradesCmds()...)
	RootCmd.AddCommand(futuresCmd)
}

type Exchange interface {
	GetBalances() (balance, pnl float64, err error)
}
