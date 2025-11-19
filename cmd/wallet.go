package cmd

import (
	"github.com/UnipayFI/binance-cli/cmd/wallet"
	"github.com/spf13/cobra"
)

var (
	walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Wallet",
	}
)

func init() {
	walletCmd.AddCommand(wallet.InitDustCmds()...)
	walletCmd.AddCommand(wallet.InitFeeCmds()...)
	walletCmd.AddCommand(wallet.InitUniversalTransferCmds()...)
	RootCmd.AddCommand(walletCmd)
}
