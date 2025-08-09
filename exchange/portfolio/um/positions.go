package um

import (
	"context"
	"errors"

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

func (c *Client) GetUMPositionRisk(symbol string) (PositionRiskList, error) {
	service := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMPositionRiskService()
	if symbol != "" {
		service.Symbol(symbol)
	}
	risks, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return risks, nil
}

func (c *Client) GetPositionSide() (bool, error) {
	service := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMPositionModeService()
	mode, err := service.Do(context.Background())
	if err != nil {
		return false, err
	}
	return mode.DualSidePosition, nil
}

func (c *Client) ChangePositionSide(dualSidePosition bool) error {
	resp, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewChangeUMPositionModeService().DualSidePosition(dualSidePosition).Do(context.Background())
	if err != nil {
		return err
	}
	if resp.Code != 200 {
		return errors.New(resp.Msg)
	}
	return nil
}
