package unified

import (
	"time"

	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2/portfolio"
)

var _ printer.TableWriter = (*AccountBalanceList)(nil)
var _ printer.TableWriter = (*PositionList)(nil)

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

type PositionList []portfolio.UMPositionV2

func (a *PositionList) Header() []string {
	return []string{"Symbol", "Initial Margin", "Maintenance Margin", "Unrealized Profit", "Position Side", "Position Amount", "Notional", "Update Time"}
}

func (a *PositionList) Row() [][]any {
	rows := [][]any{}
	for _, position := range *a {
		rows = append(rows, []any{
			position.Symbol,
			position.InitialMargin,
			position.MaintMargin,
			position.UnrealizedProfit,
			position.PositionSide,
			position.PositionAmt,
			position.Notional,
			time.UnixMilli(position.UpdateTime).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}

type OrderList []portfolio.UMAllOrdersResponse

func (a *OrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (a *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *a {
		rows = append(rows, []any{
			order.OrderID,
			order.Symbol,
			order.Side,
			order.Status,
			order.Price,
			order.OrigQty,
			order.ExecutedQty,
			time.UnixMilli(order.Time).Format("2006-01-02 15:04:05"),
			time.UnixMilli(order.UpdateTime).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}
