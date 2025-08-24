package um

import (
	"time"

	"github.com/adshao/go-binance/v2/portfolio"
)

type SymbolConfigList []*portfolio.UMSymbolConfig

func (s *SymbolConfigList) Header() []string {
	return []string{"Symbol", "Margin Type", "Is Auto Add Margin", "Leverage", "Max Notional Value"}
}

func (s *SymbolConfigList) Row() [][]any {
	rows := [][]any{}
	for _, symbol := range *s {
		rows = append(rows, []any{
			symbol.Symbol,
			symbol.MarginType,
			symbol.IsAutoAddMargin,
			symbol.Leverage,
			symbol.MaxNotionalValue,
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

type PositionRiskList []*portfolio.UMPosition

func (a *PositionRiskList) Header() []string {
	return []string{"Symbol", "Position Side", "Position Amont", "Entry Price", "Unrealized Profit", "Liquidation Price", "Leverage", "Notional", "Max Notional", "Update Time"}
}

func (a *PositionRiskList) Row() [][]any {
	rows := [][]any{}
	for _, position := range *a {
		rows = append(rows, []any{
			position.Symbol,
			position.PositionSide,
			position.PositionAmt,
			position.EntryPrice,
			position.UnrealizedProfit,
			position.LiquidationPrice,
			position.Leverage,
			position.Notional,
			position.MaxNotionalValue,
			time.UnixMilli(position.UpdateTime).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}

type IncomeHistoryList []*portfolio.Income

func (i *IncomeHistoryList) Header() []string {
	return []string{"TranID", "TradeID", "Symbol", "Income Type", "Income", "Asset", "Info", "Time"}
}

func (i *IncomeHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, income := range *i {
		rows = append(rows, []any{income.TranID, income.TradeID, income.Symbol, income.IncomeType, income.Income, income.Asset, income.Info, time.UnixMilli(income.Time).Format("2006-01-02 15:04:05")})
	}
	return rows
}

type OrderList []portfolio.UMAllOrdersResponse

func (a *OrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Type", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (a *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *a {
		rows = append(rows, []any{
			order.OrderID,
			order.Symbol,
			order.Side,
			order.Type,
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

type OpenOrderList []portfolio.UMOpenOrdersResponse

func (a *OpenOrderList) Header() []string {
	return []string{"Order ID", "Symbol", "Side", "Status", "Price", "Quantity", "Executed Quantity", "Time", "Update Time"}
}

func (a *OpenOrderList) Row() [][]any {
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
