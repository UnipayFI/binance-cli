package portfolio

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetAutoRepayStatus() (bool, error) {
	resp, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetAutoRepayFuturesStatusService().Do(context.Background())
	if err != nil {
		return false, err
	}
	return resp.AutoRepay, nil
}

func (c *Client) SetAutoRepayStatus(status bool) (bool, error) {
	_, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewChangeAutoRepayFuturesStatusService().AutoRepay(status).Do(context.Background())
	if err != nil {
		return status, err
	}
	return status, nil
}
