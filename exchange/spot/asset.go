package spot

import (
	"context"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetAssetList() (AssetBalanceList, error) {
	account, err := binance.NewClient(c.ApiKey, c.ApiSecret).NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	balances := AssetBalanceList{}
	for i := range account.Balances {
		balance := account.Balances[i]
		if !common.IsZero(balance.Free) || !common.IsZero(balance.Locked) {
			balances = append(balances, balance)
		}
	}
	return AssetBalanceList(balances), nil
}
