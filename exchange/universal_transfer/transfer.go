package universaltransfer

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2"
)

type Client struct {
	*exchange.Client
}

func (c *Client) Transfer(transferType, asset, amount, fromSymbol, toSymbol string) error {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewUserUniversalTransferService()
	service.Type(binance.UserUniversalTransferType(transferType))
	service.Asset(asset)
	service.Amount(amount)
	if fromSymbol != "" {
		service.FromSymbol(fromSymbol)
	}
	if toSymbol != "" {
		service.ToSymbol(toSymbol)
	}
	_, err := service.Do(context.Background())
	return err
}
