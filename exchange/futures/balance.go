package futures

import (
	"context"

	"github.com/UnipayFI/binance-cli/common"
	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/futures"
)

type Client struct {
	*exchange.Client
}

func (c *Client) GetBalances() (balance *BalanceList, err error) {
	balances, err := futures.NewClient(c.ApiKey, c.ApiSecret).NewGetBalanceService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	balance = &BalanceList{}
	for i := range balances {
		if common.IsZero(balances[i].Balance) && common.IsZero(balances[i].CrossWalletBalance) && common.IsZero(balances[i].CrossUnPnl) && common.IsZero(balances[i].AvailableBalance) {
			continue
		}
		*balance = append(*balance, *balances[i])
	}
	return balance, nil
}
