package unified

import (
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/unified"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	marginCmd = &cobra.Command{
		Use:   "margin",
		Short: "margin",
	}

	marginLoanCmd = &cobra.Command{
		Use:   "loan",
		Short: "loan",
		Run:   marginLoan,
	}
	marginRepayCmd = &cobra.Command{
		Use:   "repay",
		Short: "repay",
		Run:   marginRepay,
	}

	marginInterestHistoryCmd = &cobra.Command{
		Use:   "interest-history",
		Short: "interest history",
		Run:   marginInterestHistory,
	}
)

func init() {
	marginCmd.PersistentFlags().StringP("asset", "a", "", "asset")
	marginCmd.MarkFlagRequired("asset")
}

func InitMarginCmds() []*cobra.Command {
	marginCmd.AddCommand(marginLoanCmd, marginRepayCmd, marginInterestHistoryCmd)
	return []*cobra.Command{marginCmd}
}

func marginLoan(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	asset, _ := cmd.Flags().GetString("asset")
	list, err := client.GetMarginLoan(asset)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}

func marginRepay(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	asset, _ := cmd.Flags().GetString("asset")
	list, err := client.GetMarginRepay(asset)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}

func marginInterestHistory(cmd *cobra.Command, _ []string) {
	client := unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	asset, _ := cmd.Flags().GetString("asset")
	list, err := client.GetMarginInterestHistory(asset)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&list)
}
