package margin

import (
	"time"

	"github.com/adshao/go-binance/v2/portfolio"
)

type MarginLoanList []portfolio.MarginLoan

func (a *MarginLoanList) Header() []string {
	return []string{"TxID", "Asset", "Principal", "Status", "Timestamp"}
}

func (a *MarginLoanList) Row() [][]any {
	rows := [][]any{}
	for _, loan := range *a {
		rows = append(rows, []any{
			loan.TxID,
			loan.Asset,
			loan.Principal,
			loan.Status,
			time.UnixMilli(loan.Timestamp).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}

type RepayLoanList []portfolio.MarginRepay

func (a *RepayLoanList) Header() []string {
	return []string{"TxID", "Asset", "Amount", "Interest", "Principal", "Status", "Timestamp"}
}

func (a *RepayLoanList) Row() [][]any {
	rows := [][]any{}
	for _, repay := range *a {
		rows = append(rows, []any{
			repay.TxID,
			repay.Asset,
			repay.Amount,
			repay.Interest,
			repay.Principal,
			repay.Status,
			time.UnixMilli(repay.Timestamp).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}

type MarginOrderList []*portfolio.MarginOrder

func (a *MarginOrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (a *MarginOrderList) Row() [][]any {
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
			time.UnixMilli(order.TransactTime).Format("2006-01-02 15:04:05"),
			time.UnixMilli(order.UpdateTime).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}
