package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) SetUMLeverage(symbol string, leverage int) error {
	_, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewChangeUMInitialLeverageService().Leverage(leverage).Symbol(symbol).Do(context.Background())
	return err
}
