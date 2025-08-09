package margin

import (
	"fmt"
	"log"

	"github.com/UnipayFI/binance-cli/config"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/UnipayFI/binance-cli/exchange/portfolio/margin"
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/spf13/cobra"
)

var (
	loanCmd = &cobra.Command{
		Use:   "loan",
		Short: "loan",
	}

	loanExecCmd = &cobra.Command{
		Use:     "exec",
		Aliases: []string{"ex"},
		Short:   "loan exec",
		Long: `Apply for a margin loan.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Borrow`,
		Run: loanExec,
	}

	loanListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "loan list",
		Long: `Query margin loan record.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-Loan-Record`,
		Run: loanList,
	}

	repayCmd = &cobra.Command{
		Use:   "repay",
		Short: "repay loan",
		Long: `Repay for a margin loan.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Repay`,
		Run: loanRepayExec,
	}

	repayDebtCmd = &cobra.Command{
		Use:   "repay-debt",
		Short: "repay debt",
		Long: `Repay debt for a margin loan.
- The repay asset amount cannot exceed 50000 USD equivalent value for a single request.
- If 'amount' is not sent, all the asset loan will be repaid if having enough specific repay assets.
- If 'amount' is sent, only the certain amount of the asset loan will be repaid if having enough specific repay assets.
- The system will use the same asset to repay the loan first (if have) no matter whether put the asset in 'specifyRepayAssets'.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/trade/Margin-Account-Repay-Debt`,
		Run: loanRepayDebt,
	}

	repayListCmd = &cobra.Command{
		Use:   "repay-list",
		Short: "repay history",
		Long: `Query margin repay record.

Docs Link: https://developers.binance.com/docs/derivatives/portfolio-margin/account/Query-Margin-repay-Record`,
		Run: loanRepayList,
	}
)

func InitLoanCmds() []*cobra.Command {
	loanExecCmd.Flags().StringP("asset", "a", "", "asset")
	loanExecCmd.Flags().StringP("amount", "m", "", "amount")
	loanExecCmd.MarkFlagRequired("asset")
	loanExecCmd.MarkFlagRequired("amount")

	loanListCmd.Flags().StringP("asset", "a", "", "asset")
	loanListCmd.Flags().StringP("txId", "i", "", "txId")
	loanListCmd.Flags().Int64P("startTime", "t", 0, "startTime")
	loanListCmd.Flags().Int64P("endTime", "e", 0, "endTime")
	loanListCmd.Flags().IntP("current", "c", 1, "Currently querying page. Start from 1.")
	loanListCmd.Flags().IntP("size", "s", 10, "size, max: 100")
	loanListCmd.Flags().BoolP("archived", "o", false, "Set to true for archived data from 6 months ago")
	loanListCmd.MarkFlagRequired("asset")

	repayCmd.Flags().StringP("asset", "a", "", "asset")
	repayCmd.Flags().StringP("amount", "m", "", "amount")
	repayCmd.MarkFlagRequired("asset")
	repayCmd.MarkFlagRequired("amount")

	repayDebtCmd.Flags().StringP("asset", "a", "", "asset")
	repayDebtCmd.Flags().StringP("amount", "m", "", "amount")
	repayDebtCmd.Flags().StringP("specifyRepayAssets", "s", "", "Specific asset list to repay debt; Can be added in batch, separated by commas")
	repayDebtCmd.MarkFlagRequired("asset")

	repayListCmd.Flags().StringP("asset", "a", "", "asset")
	repayListCmd.Flags().StringP("txId", "i", "", "txId")
	repayListCmd.Flags().Int64P("startTime", "t", 0, "startTime")
	repayListCmd.Flags().Int64P("endTime", "e", 0, "endTime")
	repayListCmd.Flags().IntP("current", "c", 1, "Currently querying page. Start from 1.")
	repayListCmd.Flags().IntP("size", "s", 10, "size, max: 100")
	repayListCmd.Flags().BoolP("archived", "o", false, "Set to true for archived data from 6 months ago")
	repayListCmd.MarkFlagRequired("asset")

	loanCmd.AddCommand(loanExecCmd, loanListCmd, repayCmd, repayDebtCmd, repayListCmd)
	return []*cobra.Command{loanCmd}
}

func loanExec(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	amount, _ := cmd.Flags().GetString("amount")

	client := margin.NewClient(config.Config.APIKey, config.Config.APISecret)
	id, err := client.LoanExec(asset, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("loan transaction id: %d\n", id)
}

func loanList(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	txId, _ := cmd.Flags().GetString("txId")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	current, _ := cmd.Flags().GetInt("current")
	size, _ := cmd.Flags().GetInt("size")
	archived, _ := cmd.Flags().GetBool("archived")
	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	loans, err := client.LoanList(asset, txId, startTime, endTime, current, size, archived)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&loans)
}

func loanRepayExec(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	amount, _ := cmd.Flags().GetString("amount")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	id, err := client.LoanRepayExec(asset, amount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repay transaction id: %d\n", id)
}

func loanRepayDebt(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	amount, _ := cmd.Flags().GetString("amount")
	specifyRepayAssets, _ := cmd.Flags().GetString("specifyRepayAssets")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	_, err := client.LoanRepayDebtExec(asset, amount, specifyRepayAssets)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repay %s %s successfully\n", asset, amount)
}

func loanRepayList(cmd *cobra.Command, args []string) {
	asset, _ := cmd.Flags().GetString("asset")
	txId, _ := cmd.Flags().GetString("txId")
	startTime, _ := cmd.Flags().GetInt64("startTime")
	endTime, _ := cmd.Flags().GetInt64("endTime")
	current, _ := cmd.Flags().GetInt("current")
	size, _ := cmd.Flags().GetInt("size")
	archived, _ := cmd.Flags().GetBool("archived")

	client := margin.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
	repays, err := client.LoanRepayList(asset, txId, startTime, endTime, current, size, archived)
	if err != nil {
		log.Fatal(err)
	}
	printer.PrintTable(&repays)
}
