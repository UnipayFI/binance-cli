package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/futures"
	"github.com/spf13/cobra"
)

var futuresCmd = &cobra.Command{
	Use:   "futures",
	Short: "futures",
}

func init() {
	futuresCmd.AddCommand(futures.InitBalancesCmds()...)
	futuresCmd.AddCommand(futures.InitOrderCmds()...)
	futuresCmd.AddCommand(futures.InitFeeCmds()...)
	futuresCmd.AddCommand(futures.InitIncomeCmds()...)
	futuresCmd.AddCommand(futures.InitPositionsCmds()...)
	futuresCmd.AddCommand(futures.InitPositionSideCmds()...)
	RootCmd.AddCommand(futuresCmd)
}
