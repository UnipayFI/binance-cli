package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetUMIncome() (UMIncomeHistoryList, error) {
	incomeHistory, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMIncomeHistoryService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	incomeHistoryList := make(UMIncomeHistoryList, 0, len(incomeHistory))
	for _, income := range incomeHistory {
		incomeHistoryList = append(incomeHistoryList, income)
	}
	return incomeHistoryList, nil
}
