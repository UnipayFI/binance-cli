package portfolio

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) BnbTransfer(amount, transferSide string) (*portfolio.BNBTransferResponse, error) {
	return portfolio.NewClient(c.ApiKey, c.ApiSecret).NewBNBTransferService().Amount(amount).TransferSide(transferSide).Do(context.Background())
}
