package portfolio

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

func (c *Client) GetMarginInterestHistory(asset string, startTime int64, endTime int64, current int64, size int64, archived bool) (MarginInterestHistoryList, error) {
	service := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginInterestHistoryService()
	if asset != "" {
		service = service.Asset(asset)
	}
	if startTime != 0 {
		service = service.StartTime(startTime)
	}
	if endTime != 0 {
		service = service.EndTime(endTime)
	}
	if current != 0 {
		service = service.Current(current)
	}
	if size != 0 {
		service = service.Size(size)
	}
	if archived {
		service = service.Archived(archived)
	}
	interestHistory, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	interestHistoryList := make(MarginInterestHistoryList, 0, len(interestHistory.Rows))
	for _, interest := range interestHistory.Rows {
		interestHistoryList = append(interestHistoryList, &interest)
	}
	return interestHistoryList, nil
}
