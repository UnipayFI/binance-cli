package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/portfolio_pro"
	"github.com/spf13/cobra"
)

var (
	portfolioProCmd = &cobra.Command{
		Use:   "portfolio-pro",
		Short: "portfolio-pro",
	}
)

func init() {
	portfolioProCmd.AddCommand(portfolio_pro.InitBFUSDCmds()...)
	RootCmd.AddCommand(portfolioProCmd)
}
