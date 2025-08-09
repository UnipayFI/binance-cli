package margin

import (
	"context"

	"github.com/UnipayFI/binance-cli/exchange"
	"github.com/adshao/go-binance/v2/portfolio"
)

type Client struct {
	*exchange.Client
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{Client: exchange.NewClient(apiKey, apiSecret)}
}

func (c *Client) LoanExec(asset, amount string) (int64, error) {
	loan, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewMarginLoanService().Asset(asset).Amount(amount).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return loan.TranID, nil
}

func (c *Client) LoanList(asset, txId string, startTime, endTime int64, current, size int, archived bool) (MarginLoanList, error) {
	loan, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginLoanService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return MarginLoanList(loan.Rows), nil
}

func (c *Client) LoanRepayExec(asset, amount string) (int64, error) {
	repay, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewMarginRepayService().Asset(asset).Amount(amount).Do(context.Background())
	if err != nil {
		return 0, err
	}
	return repay.TranID, nil
}

func (c *Client) LoanRepayDebtExec(asset, amount, specifyRepayAssets string) (*portfolio.MarginRepayDebtResponse, error) {
	repay, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewMarginRepayDebtService().Asset(asset).Amount(amount).SpecifyRepayAssets(specifyRepayAssets).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return repay, nil
}

func (c *Client) LoanRepayList(asset, txId string, startTime, endTime int64, current, size int, archived bool) (RepayLoanList, error) {
	repay, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginRepayService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return RepayLoanList(repay.Rows), nil
}
