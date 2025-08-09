package portfolio

import (
	"github.com/UnipayFI/binance-cli/cmd/portfolio/um"
	"github.com/spf13/cobra"
)

var (
	umCmd = &cobra.Command{
		Use:   "um",
		Short: "USDⓈ-Margined Futures",
	}
)

func InitUMCmds() []*cobra.Command {
	umCmd.AddCommand(um.InitCommissionRateCmds()...)
	umCmd.AddCommand(um.InitFeeCmds()...)
	umCmd.AddCommand(um.InitIncomeCmds()...)
	umCmd.AddCommand(um.InitOrderCmds()...)
	umCmd.AddCommand(um.InitPositionsCmds()...)
	umCmd.AddCommand(um.InitSymbolConfigCmds()...)
	return []*cobra.Command{umCmd}
}
