package universaltransfer

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
