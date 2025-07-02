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
	UnifiedCmd.AddCommand(unified.InitAccountCmds()...)
	UnifiedCmd.AddCommand(unified.InitCollectionCmds()...)
	UnifiedCmd.AddCommand(unified.InitAutoRepayCmds()...)
	UnifiedCmd.AddCommand(unified.InitBalancesCmds()...)
	UnifiedCmd.AddCommand(unified.InitBnbTransferCmds()...)
	UnifiedCmd.AddCommand(unified.InitCommissionRateCmds()...)
	UnifiedCmd.AddCommand(unified.InitFeeCmds()...)
	UnifiedCmd.AddCommand(unified.InitIncomeCmds()...)
	UnifiedCmd.AddCommand(unified.InitInterestHistoryCmds()...)
	UnifiedCmd.AddCommand(unified.InitLeverageCmds()...)
	UnifiedCmd.AddCommand(unified.InitMarginCmds()...)
	UnifiedCmd.AddCommand(unified.InitOrderCmds()...)
	UnifiedCmd.AddCommand(unified.InitPositionsCmds()...)
	UnifiedCmd.AddCommand(unified.InitRepayFuturesNegativeBalanceCmds()...)

	RootCmd.AddCommand(UnifiedCmd)
}
