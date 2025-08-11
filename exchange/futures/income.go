package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetIncome(symbol string, incomeType string, startTime int64, endTime int64, page, limit int64) (IncomeHistoryList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetIncomeHistoryService()
	if symbol != "" {
		service.Symbol(symbol)
	}
	if incomeType != "" {
		service.IncomeType(incomeType)
	}
	if startTime != 0 {
		service.StartTime(startTime)
	}
	if endTime != 0 {
		service.EndTime(endTime)
	}
	if page != 0 {
		service.Page(page)
	}
	if limit != 0 {
		service.Limit(int64(limit))
	}
	income, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return IncomeHistoryList(income), nil
}
