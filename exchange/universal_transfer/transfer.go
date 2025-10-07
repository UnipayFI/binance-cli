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

func (c *Client) List(transferType string, startTime, endTime, current, size int64, fromSymbol, toSymbol string) (UniversalTransferList, error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewListUserUniversalTransferService()
	service.Type(binance.UserUniversalTransferType(transferType))
	if startTime != 0 {
		service.StartTime(startTime)
	}
	if endTime != 0 {
		service.EndTime(endTime)
	}
	if current != 0 {
		service.Current(int(current))
	}
	if size != 0 {
		service.Size(int(size))
	}
	if fromSymbol != "" {
		service.FromSymbol(fromSymbol)
	}
	if toSymbol != "" {
		service.ToSymbol(toSymbol)
	}
	resp, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return UniversalTransferList(resp.Results), nil
}
