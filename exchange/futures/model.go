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

type OrderList []*futures.Order

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Client Order ID", "Symbol", "Side", "Position Side", "Status", "Price", "Quantity", "Executed Quantity"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{order.OrderID, order.ClientOrderID, order.Symbol, order.Side, order.PositionSide, order.Status, order.Price, order.OrigQuantity, order.ExecutedQuantity})
	}
	return rows
}
