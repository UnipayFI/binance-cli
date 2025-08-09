package futures

import (
	"context"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetSymbolConfig(symbol string) (SymbolConfigList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetSymbolConfigService()
	service.Symbol(symbol)
	symbolConfig, err := service.Do(context.Background())
	if err != nil {
		return SymbolConfigList{}, err
	}
	return symbolConfig, nil
}

func (c *Client) SetLeverage(symbol string, leverage int) error {
	_, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewChangeLeverageService().Symbol(symbol).Leverage(leverage).Do(context.Background())
	return err
}

func (c *Client) SetMarginType(symbol string, marginType futures.MarginType) error {
	return futures.NewClient(c.ApiKey, c.ApiSecret).NewChangeMarginTypeService().Symbol(symbol).MarginType(marginType).Do(context.Background())
}
