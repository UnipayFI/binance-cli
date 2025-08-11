package um

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) FeeBurnStatus() (*portfolio.UMFeeBurnStatusResponse, error) {
	feeBurn, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewUMFeeBurnStatusService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return feeBurn, nil
}

func (c *Client) SetFeeBurnStatus(status bool) error {
	_, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewUMFeeBurnService().FeeBurn(status).Do(context.Background())
	return err
}
