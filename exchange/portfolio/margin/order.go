package margin

import (
	"context"

	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetOrderList(symbol string, orderID, startTime, endTime int64, limit int) (MarginOrderList, error) {
	orderService := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginAllOrdersService().Symbol(symbol).OrderID(orderID).StartTime(startTime).EndTime(endTime).Limit(limit)
	orders, err := orderService.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return MarginOrderList(orders), nil
}

func (c *Client) GetOpenOrderList(symbol string) (MarginOrderList, error) {
	orderService := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewGetMarginOpenOrdersService().Symbol(symbol)
	orders, err := orderService.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return MarginOrderList(orders), nil
}

func (c *Client) CreateOrder(params map[string]string) (*portfolio.MarginOrder, error) {
	orderService := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewMarginOrderService()
	if params["symbol"] != "" {
		orderService.Symbol(params["symbol"])
	}
	if params["side"] != "" {
		orderService.Side(portfolio.SideType(params["side"]))
	}
	if params["type"] != "" {
		orderService.Type(portfolio.OrderType(params["type"]))
	}
	if params["quantity"] != "" {
		orderService.Quantity(params["quantity"])
	}
	if params["quoteOrderQty"] != "" {
		orderService.QuoteOrderQty(params["quoteOrderQty"])
	}
	if params["price"] != "" {
		orderService.Price(params["price"])
	}
	if params["stopPrice"] != "" {
		orderService.StopPrice(params["stopPrice"])
	}
	if params["newClientOrderId"] != "" {
		orderService.NewClientOrderID(params["newClientOrderId"])
	}
	if params["newOrderRespType"] != "" {
		orderService.NewOrderRespType(portfolio.NewOrderRespType(params["newOrderRespType"]))
	}
	if params["icebergQty"] != "" {
		orderService.IcebergQty(params["icebergQty"])
	}
	if params["sideEffectType"] != "" {
		orderService.SideEffectType(params["sideEffectType"])
	}
	if params["timeInForce"] != "" {
		orderService.TimeInForce(portfolio.TimeInForceType(params["timeInForce"]))
	}
	if params["selfTradePreventionMode"] != "" {
		orderService.SelfTradePreventionMode(portfolio.SelfTradePreventionMode(params["selfTradePreventionMode"]))
	}
	if params["autoRepayAtCancel"] != "" {
		orderService.AutoRepayAtCancel(params["autoRepayAtCancel"] == "true")
	}
	order, err := orderService.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) error {
	orderService := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewMarginCancelOrderService().Symbol(symbol)
	if orderID != 0 {
		orderService.OrderID(orderID)
	}
	if clientOrderID != "" {
		orderService.OrigClientOrderID(clientOrderID)
	}
	_, err := orderService.Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
