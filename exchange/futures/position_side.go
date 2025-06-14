package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetPositionSide() (bool, error) {
	positionSide, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetPositionModeService().Do(context.Background())
	if err != nil {
		return false, err
	}
	return positionSide.DualSidePosition, nil
}

func (c *Client) ChangePositionSide(dualSide bool) error {
	err := futures.NewClient(c.ApiKey, c.ApiSecret).NewChangePositionModeService().DualSide(dualSide).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
