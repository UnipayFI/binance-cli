package futures

import (
	"fmt"
	"strconv"
	"strings"
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

type AccountConfigList []*futures.AccountConfig

func (a *AccountConfigList) Header() []string {
	return []string{"Fee Tier", "Can Trade", "Can Deposit", "Can Withdraw", "Dual Side Position", "Multi Assets Margin", "Trade Group ID"}
}

func (a *AccountConfigList) Row() [][]any {
	rows := [][]any{}
	for _, config := range *a {
		rows = append(rows, []any{config.FeeTier, config.CanTrade, config.CanDeposit, config.CanWithdraw, config.DualSidePosition, config.MultiAssetsMargin, config.TradeGroupId})
	}
	return rows
}

type ForceOrderList []*futures.UserLiquidationOrder

func (f *ForceOrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Position Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (f *ForceOrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *f {
		rows = append(rows, []any{order.OrderId, order.Symbol, order.Side, order.PositionSide, order.Status, order.Price, order.OrigQuantity, order.ExecutedQuantity, time.UnixMilli(order.Time).Format("2006-01-02 15:04:05"), time.UnixMilli(order.UpdateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type OrderList []*futures.Order

func (o *OrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Type", "Position Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{order.OrderID, order.Symbol, order.Side, order.Type, order.PositionSide, order.Status, order.Price, order.OrigQuantity, order.ExecutedQuantity, time.UnixMilli(order.Time).Format("2006-01-02 15:04:05"), time.UnixMilli(order.UpdateTime).Format("2006-01-02 15:04:05")})
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

type TradeList []*futures.AccountTrade

func (t *TradeList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Position Side", "Price", "Quantity", "Quote Quantity", "Realized Pnl", "Time"}
}

func (t *TradeList) Row() [][]any {
	rows := [][]any{}
	for _, trade := range *t {
		rows = append(rows, []any{trade.OrderID, trade.Symbol, trade.Side, trade.PositionSide, trade.Price, trade.Quantity, trade.QuoteQuantity, trade.RealizedPnl, time.UnixMilli(trade.Time).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type SymbolConfigList []*futures.SymbolConfig

func (s *SymbolConfigList) Header() []string {
	return []string{"Symbol", "Margin Type", "Is Auto Add Margin", "Leverage", "Max Notional Value"}
}

func (s *SymbolConfigList) Row() [][]any {
	rows := [][]any{}
	for _, symbolConfig := range *s {
		rows = append(rows, []any{symbolConfig.Symbol, symbolConfig.MarginType, symbolConfig.IsAutoAddMargin, symbolConfig.Leverage, symbolConfig.MaxNotionalValue})
	}
	return rows
}

type PositionRiskList []*futures.PositionRiskV3

func (p *PositionRiskList) Header() []string {
	return []string{"Symbol", "Position Side", "Position Amount", "Notional", "Entry Price", "Mark Price", "Unrealized Profit", "Liquidation Price", "Update Time"}
}

func (p *PositionRiskList) Row() [][]any {
	rows := [][]any{}
	for _, risk := range *p {
		asset := strings.TrimRight(risk.Symbol, risk.MarginAsset)
		rows = append(rows, []any{risk.Symbol, risk.PositionSide, risk.PositionAmt + " " + asset, risk.Notional, risk.EntryPrice, risk.MarkPrice, risk.UnRealizedProfit, risk.LiquidationPrice, time.UnixMilli(risk.UpdateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type CommissionRateList []*futures.CommissionRate

func (c *CommissionRateList) Header() []string {
	return []string{"Symbol", "Maker Commission Rate", "Taker Commission Rate"}
}

func (c *CommissionRateList) Row() [][]any {
	rows := [][]any{}
	for _, rate := range *c {
		maker, _ := strconv.ParseFloat(rate.MakerCommissionRate, 64)
		taker, _ := strconv.ParseFloat(rate.TakerCommissionRate, 64)
		rows = append(rows, []any{rate.Symbol, fmt.Sprintf("%.4f%%", maker*100), fmt.Sprintf("%.4f%%", taker*100)})
	}
	return rows
}
