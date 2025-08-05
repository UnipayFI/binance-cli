package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/spot"
	"github.com/spf13/cobra"
)

var (
	SpotCmd = &cobra.Command{
		Use:   "spot",
		Short: "spot",
	}
)

func init() {
	SpotCmd.AddCommand(spot.InitAccountCmds()...)
	SpotCmd.AddCommand(spot.InitAssetCmds()...)
	SpotCmd.AddCommand(spot.InitOrderCmds()...)
	RootCmd.AddCommand(SpotCmd)
}
