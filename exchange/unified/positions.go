package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetUMPositions() (PositionList, error) {
	account, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMAccountDetailV2Service().Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := PositionList{}
	for _, position := range account.Positions {
		list = append(list, *position)
	}
	return list, nil
}
