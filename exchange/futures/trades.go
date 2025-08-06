package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetTrades(symbol string, orderId int64, startTime, endTime int64, fromId int64, limit int) (TradeList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewListAccountTradeService()
	if symbol != "" {
		service.Symbol(symbol)
	}
	if orderId != 0 {
		service.OrderID(orderId)
	}
	if startTime != 0 {
		service.StartTime(startTime)
	}
	if endTime != 0 {
		service.EndTime(endTime)
	}
	if fromId != 0 {
		service.FromID(fromId)
	}
	if limit != 0 {
		service.Limit(limit)
	}
	trades, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return trades, nil
}
