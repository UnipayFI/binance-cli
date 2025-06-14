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
	futuresCmd.AddCommand(futures.InitBalanceCmds()...)
	futuresCmd.AddCommand(futures.InitOrderCmds()...)
	futuresCmd.AddCommand(futures.InitFeeCmds()...)
	RootCmd.AddCommand(futuresCmd)
}
