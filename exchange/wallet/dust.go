package wallet

import (
	"context"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func (c *Client) ShowDust(accountType string) (list DustDetailList, totalBTC, totalBNB, percentage string, err error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewListDustService()
	// service.AccountType(accountType)
	response, err := service.Do(context.Background())
	if err != nil {
		return nil, "", "", "", err
	}
	return response.Details, response.TotalTransferBtc, response.TotalTransferBNB, response.DribbletPercentage, nil
}

func (c *Client) ConvertDust(asset string, accountType string) (list DustTransferResultList, totalServiceCharge, totalTransfered string, err error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewDustTransferService()
	service.Asset(strings.Split(asset, ","))
	// service.AccountType(accountType)
	response, err := service.Do(context.Background())
	if err != nil {
		return nil, "", "", err
	}
	return response.TransferResult, response.TotalServiceCharge, response.TotalTransfered, nil
}

func (c *Client) HistoryDust(startTime, endTime int64) (list DustHistoryList, err error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewListDustLogService()
	if startTime != 0 {
		service.StartTime(startTime)
	}
	if endTime != 0 {
		service.EndTime(endTime)
	}
	response, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list = make(DustHistoryList, 0, int(response.Total))
	for _, dust := range response.UserAssetDribblets {
		for _, detail := range dust.UserAssetDribbletDetails {
			list = append(list, DustHistory{
				FromAsset:           detail.FromAsset,
				Amount:              detail.Amount,
				TransID:             int64(detail.TransID),
				ServiceChargeAmount: detail.ServiceChargeAmount,
				TransferedAmount:    detail.TransferedAmount,
				OperateTime:         detail.OperateTime,
			})
		}
	}
	return list, nil
}
