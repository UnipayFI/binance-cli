package futures

import (
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2/futures"
)

var _ printer.TableWriter = (*BalanceList)(nil)

type BalanceList []futures.Balance

func (a *BalanceList) Header() []string {
	return []string{"Asset", "Balance", "Cross Wallet Balance", "Cross Un Pnl", "Available Balance", "Max Withdraw Amount"}
}

func (a *BalanceList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.Asset, asset.Balance, asset.CrossWalletBalance, asset.CrossUnPnl, asset.AvailableBalance, asset.MaxWithdrawAmount})
	}
	return rows
}
