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
	return []string{"Symbol", "Position Amont", "Entry Price", "Unrealized Profit", "Liquidation Price", "Leverage", "Max Notional", "Position Side", "Notional", "Update Time"}
}

func (a *PositionRiskList) Row() [][]any {
	rows := [][]any{}
	for _, position := range *a {
		rows = append(rows, []any{
			position.Symbol,
			position.PositionAmt,
			position.EntryPrice,
			position.UnrealizedProfit,
			position.LiquidationPrice,
			position.Leverage,
			position.MaxNotional,
			position.MaxNotionalValue,
			position.PositionSide,
			position.Notional,
			time.UnixMilli(position.UpdateTime).Format("2006-01-02 15:04:05"),
		})
	}
	return rows
}
