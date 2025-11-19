package wallet

import (
	"time"

	"github.com/UnipayFI/binance-cli/printer"
	"github.com/adshao/go-binance/v2"
)

var _ printer.TableWriter = (*UniversalTransferList)(nil)

type UniversalTransferList []*binance.UserUniversalTransfer

func (a *UniversalTransferList) Header() []string {
	return []string{"Asset", "Amount", "Type", "Status", "TranId", "Timestamp"}
}

func (a *UniversalTransferList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.Asset, asset.Amount, asset.Type, asset.Status, asset.TranId, time.UnixMilli(asset.Timestamp).Format("2006-01-02 15:04:05")})
	}
	return rows
}

var _ printer.TableWriter = (*DustDetailList)(nil)

type DustDetailList []binance.ListDustDetail

func (a *DustDetailList) Header() []string {
	return []string{"Asset", "Amount", "toBTC", "toBNB", "toBNBOffExchange", "Commission fee"}
}

func (a *DustDetailList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.Asset, asset.AmountFree, asset.ToBTC, asset.ToBNB, asset.ToBNBOffExchange, asset.Exchange})
	}
	return rows
}

var _ printer.TableWriter = (*DustTransferResultList)(nil)

type DustTransferResultList []*binance.DustTransferResult

func (a *DustTransferResultList) Header() []string {
	return []string{"From Asset", "Amount", "Transfer ID", "Service Charge Amount", "Transferred Amount", "Operate Time"}
}

func (a *DustTransferResultList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.FromAsset, asset.Amount, asset.TranID, asset.ServiceChargeAmount, asset.TransferedAmount, time.UnixMilli(asset.OperateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}

var _ printer.TableWriter = (*DustHistoryList)(nil)

type DustHistoryList []DustHistory

type DustHistory struct {
	FromAsset           string
	Amount              string
	TransID             int64
	ServiceChargeAmount string
	TransferedAmount    string
	OperateTime         int64
}

func (a *DustHistoryList) Header() []string {
	return []string{"From Asset", "Amount", "Transfer ID", "Service Charge Amount", "Transferred Amount", "Operate Time"}
}

func (a *DustHistoryList) Row() [][]any {
	rows := [][]any{}
	for _, asset := range *a {
		rows = append(rows, []any{asset.FromAsset, asset.Amount, asset.TransID, asset.ServiceChargeAmount, asset.TransferedAmount, time.UnixMilli(asset.OperateTime).Format("2006-01-02 15:04:05")})
	}
	return rows
}
