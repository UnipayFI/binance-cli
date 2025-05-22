package spot

import (
	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2"
)

var _ printer.TableWriter = (*AssetBalanceList)(nil)

type AssetBalanceList []binance.Balance

func (a *AssetBalanceList) Header() []string {
	return []string{"Asset", "Free", "Locked"}
}

func (a *AssetBalanceList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.Asset, asset.Free, asset.Locked})
	}
	return rows
}

type OrderList []*binance.Order

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Client Order ID", "Symbol", "Side", "Status", "Price", "Quantity", "Executed Quantity"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{order.OrderID, order.ClientOrderID, order.Symbol, order.Side, order.Status, order.Price, order.OrigQuantity, order.ExecutedQuantity})
	}
	return rows
}
