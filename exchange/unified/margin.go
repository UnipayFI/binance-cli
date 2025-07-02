package unified

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetMarginLoan(asset string) (MarginLoanList, error) {
	loans, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginLoanService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	loanList := make(MarginLoanList, 0, len(loans.Rows))
	for _, loan := range loans.Rows {
		loanList = append(loanList, &loan)
	}
	return loanList, nil
}

func (c *Client) GetMarginRepay(asset string) (RepayLoanList, error) {
	repays, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginRepayService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	repayList := make(RepayLoanList, 0, len(repays.Rows))
	for _, repay := range repays.Rows {
		repayList = append(repayList, &repay)
	}
	return repayList, nil
}

func (c *Client) GetMarginInterestHistory(asset string) (InterestHistoryList, error) {
	interestHistory, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginInterestHistoryService().Asset(asset).Do(context.Background())
	if err != nil {
		return nil, err
	}
	interestHistoryList := make(InterestHistoryList, 0, len(interestHistory.Rows))
	for _, interest := range interestHistory.Rows {
		interestHistoryList = append(interestHistoryList, &interest)
	}
	return interestHistoryList, nil
}
