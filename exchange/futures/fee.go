package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) FeeBurnStatus() (*futures.FeeBurn, error) {
	feeBurn, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetFeeBurnService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return feeBurn, nil
}

func (c *Client) FeeBurnStatusChange(status bool) error {
	bnbService := futures.NewClient(c.ApiKey, c.ApiSecret).NewFeeBurnService()
	if status {
		bnbService.Enable()
	} else {
		bnbService.Disable()
	}

	return bnbService.Do(context.Background())
}
