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

	SpotAssetsCmd = &cobra.Command{
		Use:   "assets",
		Short: "assets",
	}

	SpotOrdersCmd = &cobra.Command{
		Use:   "orders",
		Short: "orders",
	}
)

func init() {
	SpotCmd.AddCommand(spot.InitAssetCmds()...)
	SpotCmd.AddCommand(spot.InitOrderCmds()...)
	RootCmd.AddCommand(SpotCmd)
}
