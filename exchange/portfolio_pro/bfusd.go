package unified_pro

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/portfolio_pro"
)

type Client struct {
	*exchange.Client
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{Client: exchange.NewClient(apiKey, apiSecret)}
}

func (c *Client) Mint(fromAsset, targetAsset, amount string) (*portfolio_pro.MintBFUSDResponse, error) {
	client := portfolio_pro.NewClient(c.Client.ApiKey, c.Client.ApiSecret)
	service := client.NewMintBFUSDService().FromAsset(fromAsset).Amount(amount)
	response, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Redeem(fromAsset, targetAsset, amount string) (*portfolio_pro.RedeemBFUSDResponse, error) {
	client := portfolio_pro.NewClient(c.Client.ApiKey, c.Client.ApiSecret)
	service := client.NewRedeemBFUSDService().TargetAsset(targetAsset).Amount(amount)
	response, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return response, nil
}
