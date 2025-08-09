package spot

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetAssetList(omitZeroBalances bool) (AssetBalanceList, error) {
	account, err := binance.NewClient(c.ApiKey, c.ApiSecret).NewGetAccountService().OmitZeroBalances(omitZeroBalances).Do(context.Background())
	if err != nil {
		return nil, err
	}
	balances := AssetBalanceList{}
	for i := range account.Balances {
		balances = append(balances, account.Balances[i])
	}
	return AssetBalanceList(balances), nil
}
