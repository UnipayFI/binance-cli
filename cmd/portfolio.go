package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/portfolio"
	"github.com/spf13/cobra"
)

var (
	portfolioCmd = &cobra.Command{
		Use:   "portfolio",
		Short: "Portfolio",
	}
)

func init() {
	portfolioCmd.AddCommand(portfolio.InitAccountCmds()...)
	portfolioCmd.AddCommand(portfolio.InitCollectionCmds()...)
	portfolioCmd.AddCommand(portfolio.InitAutoRepayCmds()...)
	portfolioCmd.AddCommand(portfolio.InitBalancesCmds()...)
	portfolioCmd.AddCommand(portfolio.InitBnbTransferCmds()...)
	portfolioCmd.AddCommand(portfolio.InitMarginCmds()...)
	portfolioCmd.AddCommand(portfolio.InitUMCmds()...)

	RootCmd.AddCommand(portfolioCmd)
}
