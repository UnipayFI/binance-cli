package futures

import (
	"context"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetPositions() (PositionList, error) {
	account, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	positions := PositionList{}
	for _, position := range account.Positions {
		if !common.IsZero(position.PositionAmt) {
			positions = append(positions, position)
		}
	}
	return positions, nil
}
