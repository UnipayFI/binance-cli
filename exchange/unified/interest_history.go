package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetInterestHistory(asset string) (InterestHistoryList, error) {
	interestHistory, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetNegativeBalanceInterestHistoryService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return InterestHistoryList(interestHistory), nil
}
