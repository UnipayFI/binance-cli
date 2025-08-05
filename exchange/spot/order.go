package spot

import (
	"context"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func (c *Client) GetOrderList(symbol string, orderID, start, end int64, limit int) (OrderList, error) {
	service := binance.NewClient(c.ApiKey, c.ApiSecret).NewListOrdersService().Symbol(symbol)
	if orderID != 0 {
		service.OrderID(orderID)
	}
	if start != 0 {
		service.StartTime(start)
	}
	if end != 0 {
		service.EndTime(end)
	}
	if limit != 0 {
		service.Limit(limit)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return OrderList(orders), nil
}

func (c *Client) CreateOrder(params map[string]string) (*binance.CreateOrderResponse, error) {
	sideType := binance.SideType(strings.ToUpper(params["side"]))
	t := binance.OrderType(strings.ToUpper(params["type"]))
	orderService := binance.NewClient(c.ApiKey, c.ApiSecret).NewCreateOrderService().Symbol(params["symbol"]).Side(sideType).Type(t)

	if params["quantity"] != "" {
		orderService.Quantity(params["quantity"])
	}
	if params["quoteOrderQty"] != "" {
		orderService.QuoteOrderQty(params["quoteOrderQty"])
	}
	if params["timeInForce"] != "" {
		orderService.TimeInForce(binance.TimeInForceType(strings.ToUpper(params["timeInForce"])))
	}
	if params["price"] != "" {
		orderService.Price(params["price"])
	}
	if params["newClientOrderID"] != "" {
		orderService.NewClientOrderID(params["newClientOrderID"])
	}
	if params["stopPrice"] != "" {
		orderService.StopPrice(params["stopPrice"])
	}
	if params["trailingDelta"] != "" {
		orderService.TrailingDelta(params["trailingDelta"])
	}
	if params["icebergQuantity"] != "" {
		orderService.IcebergQuantity(params["icebergQuantity"])
	}
	if params["newOrderRespType"] != "" {
		orderService.NewOrderRespType(binance.NewOrderRespType(strings.ToUpper(params["newOrderRespType"])))
	}
	if params["selfTradePreventionMode"] != "" {
		orderService.SelfTradePreventionMode(binance.SelfTradePreventionMode(strings.ToUpper(params["selfTradePreventionMode"])))
	}

	return orderService.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) error {
	orderService := binance.NewClient(c.ApiKey, c.ApiSecret).NewCancelOrderService().Symbol(symbol)
	if orderID != 0 {
		orderService.OrderID(orderID)
	}
	if clientOrderID != "" {
		orderService.OrigClientOrderID(clientOrderID)
	}
	_, err := orderService.Do(context.Background())
	return err
}
