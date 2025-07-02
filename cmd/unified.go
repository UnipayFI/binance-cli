package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/unified"
	"github.com/spf13/cobra"
)

var (
	UnifiedCmd = &cobra.Command{
		Use:   "unified",
		Short: "unified",
	}
)

func init() {
	UnifiedCmd.AddCommand(unified.InitBalancesCmds()...)
	UnifiedCmd.AddCommand(unified.InitPositionsCmds()...)
	UnifiedCmd.AddCommand(unified.InitOrderCmds()...)
	UnifiedCmd.AddCommand(unified.InitLeverageCmds()...)
	UnifiedCmd.AddCommand(unified.InitFeeCmds()...)
	UnifiedCmd.AddCommand(unified.InitAutoCollectionCmds()...)
	RootCmd.AddCommand(UnifiedCmd)
}
