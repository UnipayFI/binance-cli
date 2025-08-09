package um

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/portfolio"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetUMSymbolConfig(symbol string) (SymbolConfigList, error) {
	symbolConfig, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetUMSymbolConfigService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return SymbolConfigList{}, err
	}
	return SymbolConfigList(symbolConfig), nil
}

func (c *Client) SetUMLeverage(symbol string, leverage int) (*portfolio.UMLeverage, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewChangeUMInitialLeverageService().Leverage(leverage).Symbol(symbol).Do(context.Background())
}
