package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetCommissionRate(symbol string) (CommissionRateList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewCommissionRateService().Symbol(symbol)
	commissionRate, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return CommissionRateList{commissionRate}, nil
}
