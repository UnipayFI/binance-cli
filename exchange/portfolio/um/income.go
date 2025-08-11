package um

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetIncome(symbol, incomeType string, startTime, endTime int64, limit, page int) (IncomeHistoryList, error) {
	incomeHistory := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMIncomeHistoryService()
	if symbol != "" {
		incomeHistory.Symbol(symbol)
	}
	if incomeType != "" {
		incomeHistory.IncomeType(incomeType)
	}
	if startTime != 0 {
		incomeHistory.StartTime(startTime)
	}
	if endTime != 0 {
		incomeHistory.EndTime(endTime)
	}
	if limit != 0 {
		incomeHistory.Limit(limit)
	}
	if page != 0 {
		incomeHistory.Page(page)
	}
	incomeHistoryList, err := incomeHistory.Do(context.Background())
	if err != nil {
		return nil, err
	}
	incomeHistoryList = IncomeHistoryList(incomeHistoryList)
	return incomeHistoryList, nil
}
