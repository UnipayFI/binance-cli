package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) SetLeverage(symbol string, leverage int) error {
	_, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewChangeLeverageService().Symbol(symbol).Leverage(leverage).Do(context.Background())
	return err
}
