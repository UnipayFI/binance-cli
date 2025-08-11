package portfolio

import (
	"time"

	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2/portfolio"
)

var _ printer.TableWriter = (*AccountBalanceList)(nil)

type AccountInfo portfolio.Account

func (a *AccountInfo) Header() []string {
	return []string{"UniMMR", "Account Equity", "Actual Equity", "Account Initial Margin", "Account Maintenance Margin", "Account Status", "Virtual Max Withdraw Amount", "Total Available Balance", "Total Margin Open Loss", "Update Time"}
}

func (a *AccountInfo) Row() [][]any {
	return [][]any{
		{a.UniMMR, a.AccountEquity, a.ActualEquity, a.AccountInitialMargin, a.AccountMaintMargin, a.AccountStatus, a.VirtualMaxWithdrawAmount, a.TotalAvailableBalance, a.TotalMarginOpenLoss, time.UnixMilli(a.UpdateTime).Format("2006-01-02 15:04:05")},
	}
}

type AccountBalanceList []portfolio.Balance

func (a *AccountBalanceList) Header() []string {
	return []string{"Asset", "Total Wallet Balance", "Cross Margin Asset", "Cross Margin Borrowed", "Cross Margin Free", "Cross Margin Interest", "Cross Margin Locked", "USDT Wallet Balance", "USDT Unrealized PNL", "Coin Wallet Balance", "Coin Unrealized PNL"}
}

func (a *AccountBalanceList) Row() [][]any {
	rows := [][]any{}
	for _, balance := range *a {
		rows = append(rows, []any{
			balance.Asset,
			balance.TotalWalletBalance,
			balance.CrossMarginAsset,
			balance.CrossMarginBorrowed,
			balance.CrossMarginFree,
			balance.CrossMarginInterest,
			balance.CrossMarginLocked,
			balance.UMWalletBalance,
			balance.UMUnrealizedPNL,
			balance.CMWalletBalance,
			balance.CMUnrealizedPNL,
		})
	}
	return rows
}

type MarginLoanList []*portfolio.MarginLoan

func (m *MarginLoanList) Header() []string {
	return []string{"TxID", "Asset", "Principal", "Status", "Timestamp"}
}

func (m *MarginLoanList) Row() [][]any {
	rows := [][]any{}
	for _, loan := range *m {
		rows = append(rows, []any{loan.TxID, loan.Asset, loan.Principal, loan.Status, time.UnixMilli(loan.Timestamp).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type RepayLoanList []*portfolio.MarginRepay

func (r *RepayLoanList) Header() []string {
	return []string{"TxID", "Asset", "Status", "Amount", "Interest", "Principal", "Timestamp"}
}

func (r *RepayLoanList) Row() [][]any {
	rows := [][]any{}
	for _, repay := range *r {
		rows = append(rows, []any{repay.TxID, repay.Asset, repay.Status, repay.Amount, repay.Interest, repay.Principal, time.UnixMilli(repay.Timestamp).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type MarginInterestHistoryList []*portfolio.MarginInterest

func (i *MarginInterestHistoryList) Header() []string {
	return []string{"TxID", "Asset", "Raw Asset", "Type", "Principal", "Interest", "Interest Rate", "Interest Accured Time"}
}

func (i *MarginInterestHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, interest := range *i {
		rows = append(rows, []any{interest.TxID, interest.Asset, interest.RawAsset, interest.Type, interest.Principal, interest.Interest, interest.InterestRate, time.UnixMilli(interest.InterestAccuredTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type InterestHistoryList []*portfolio.NegativeBalanceInterest

func (i *InterestHistoryList) Header() []string {
	return []string{"Asset", "Interest", "Interest Rate", "Principal", "Interest Accured Time"}
}

func (i *InterestHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, interest := range *i {
		rows = append(rows, []any{interest.Asset, interest.Interest, interest.InterestRate, interest.Principal, time.UnixMilli(interest.InterestAccuredTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}
