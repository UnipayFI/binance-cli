package unified

import (
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2/portfolio"
)

var _ printer.TableWriter = (*AccountBalanceList)(nil)

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
