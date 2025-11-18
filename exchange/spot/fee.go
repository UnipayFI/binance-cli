package spot

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

func (c *Client) GetFeeBurnStatus() (binance.BNBBurn, error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewGetBNBBurnService()
	response, err := service.Do(context.Background())
	if err != nil {
		return binance.BNBBurn{}, err
	}
	return *response, nil
}

func (c *Client) SetFeeBurnStatus(spotBNBBurn, interestBNBBurn bool) (binance.BNBBurn, error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewToggleBNBBurnService()
	response, err := service.SpotBNBBurn(spotBNBBurn).InterestBNBBurn(interestBNBBurn).Do(context.Background())
	if err != nil {
		return binance.BNBBurn{}, err
	}
	return *response, nil
}
