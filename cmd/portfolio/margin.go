package portfolio

import (
	"github.com/UnipayFI/binance-cli/cmd/portfolio/margin"
	"github.com/spf13/cobra"
)

var (
	marginCmd = &cobra.Command{
		Use:   "margin",
		Short: "Margin Trading",
	}
)

func InitMarginCmds() []*cobra.Command {
	marginCmd.AddCommand(margin.InitInterestHistoryCmds()...)
	marginCmd.AddCommand(margin.InitLoanCmds()...)
	marginCmd.AddCommand(margin.InitOrderCmds()...)
	return []*cobra.Command{marginCmd}
}
