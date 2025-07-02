package futures

import (
	"time"

	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2/futures"
)

var _ printer.TableWriter = (*BalanceList)(nil)
var _ printer.TableWriter = (*OrderList)(nil)

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

type OrderList []*futures.Order

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Position Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{order.OrderID, order.Symbol, order.Side, order.PositionSide, order.Status, order.Price, order.OrigQuantity, order.ExecutedQuantity, time.UnixMilli(order.Time).Format("2006-01-02 15:04:05"), time.UnixMilli(order.UpdateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type PositionList []*futures.AccountPosition

func (p *PositionList) Header() []string {
	return []string{"Symbol", "Position Side", "Position Amount", "Entry Price", "Unrealized Profit", "Leverage", "Update Time"}
}

func (p *PositionList) Row() [][]any {
	rows := [][]any{}
	for _, position := range *p {
		rows = append(rows, []any{position.Symbol, position.PositionSide, position.PositionAmt, position.EntryPrice, position.UnrealizedProfit, position.Leverage, time.UnixMilli(position.UpdateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type IncomeHistoryList []*futures.IncomeHistory

func (i *IncomeHistoryList) Header() []string {
	return []string{"Asset", "Income", "Income Type", "Info", "Symbol", "Time", "Tran ID", "Trade ID"}
}

func (i *IncomeHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, income := range *i {
		rows = append(rows, []any{income.Asset, income.Income, income.IncomeType, income.Info, income.Symbol, time.UnixMilli(income.Time).Format("2006-01-02 15:04:05"), income.TranID, income.TradeID})
	}
	return rows
}
