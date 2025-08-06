package futures

import (
	"context"
	"fmt"
	"slices"

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

func (c *Client) GetPositionRisk(symbol string) (PositionRiskList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetPositionRiskV3Service()
	if symbol != "" {
		service.Symbol(symbol)
	}
	risks, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	risks = slices.DeleteFunc(risks, func(risk *futures.PositionRiskV3) bool {
		return common.IsZero(risk.PositionAmt)
	})
	return risks, nil
}

func (c *Client) ModifyPositionMargin(symbol, positionSide string, amount float64, typ int) error {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewUpdatePositionMarginService().Symbol(symbol).Amount(fmt.Sprintf("%.6f", amount)).Type(typ)
	if positionSide != "" {
		service.PositionSide(futures.PositionSideType(positionSide))
	}
	return service.Do(context.Background())
}
