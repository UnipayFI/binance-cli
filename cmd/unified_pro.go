package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/unified_pro"
	"github.com/spf13/cobra"
)

var (
	UnifiedProCmd = &cobra.Command{
		Use:   "unified-pro",
		Short: "unified-pro",
	}
)

func init() {
	UnifiedProCmd.AddCommand(unified_pro.InitOrderCmds()...)
	RootCmd.AddCommand(UnifiedProCmd)
}
