package spot

import (
	"fmt"
	"time"

	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2"
)

var _ printer.TableWriter = (*Account)(nil)

type Account struct {
	binance.Account
}

func (a *Account) Header() []string {
	return []string{"UID", "Commission Rates", "Can Trade", "Can Withdraw", "Can Deposit", "Account Type", "Permissions", "Update Time"}
}

func (account *Account) Row() [][]any {
	a := account.Account
	rates := fmt.Sprintf("Maker: %s\nTaker: %s\nBuyer: %s\nSeller: %s", a.CommissionRates.Maker, a.CommissionRates.Taker, a.CommissionRates.Buyer, a.CommissionRates.Seller)
	return [][]any{
		{a.UID, rates, a.CanTrade, a.CanWithdraw, a.CanDeposit, a.AccountType, a.Permissions, time.UnixMilli(int64(a.UpdateTime)).Format("2006-01-02 15:04:05")},
	}
}

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

type DividendList []binance.DividendResponse

func (d *DividendList) Header() []string {
	return []string{"ID", "Asset", "Amount", "Dividend Time", "Info", "Transaction ID"}
}

func (d *DividendList) Row() [][]any {
	rows := [][]any{}
	for _, dividend := range *d {
		rows = append(rows, []any{dividend.ID, dividend.Asset, dividend.Amount, time.UnixMilli(dividend.Time).Format("2006-01-02 15:04:05"), dividend.Info, dividend.TranID})
	}
	return rows
}
