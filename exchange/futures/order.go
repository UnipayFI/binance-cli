package futures

import (
	"context"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
)

func (c *Client) GetOrderHistory(symbol string, limit int, start, end int64, orderID int64) (OrderList, error) {
	service := futures.NewClient(c.ApiKey, c.ApiSecret).NewListOrdersService()
	if symbol != "" {
		service.Symbol(symbol)
	}
	if limit != 0 {
		service.Limit(limit)
	}
	if orderID != 0 {
		service.OrderID(orderID)
	}
	if start != 0 {
		service.StartTime(start)
	}
	if end != 0 {
		service.EndTime(end)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) CreateOrder(params map[string]string) (*futures.CreateOrderResponse, error) {
	sideType := futures.SideType(strings.ToUpper(params["side"]))
	t := futures.OrderType(strings.ToUpper(params["type"]))
	orderService := futures.NewClient(c.ApiKey, c.ApiSecret).NewCreateOrderService().Symbol(params["symbol"]).Side(sideType).Type(t)

	if params["positionSide"] != "" {
		orderService.PositionSide(futures.PositionSideType(params["positionSide"]))
	}
	if params["quantity"] != "" {
		orderService.Quantity(params["quantity"])
	}
	if params["reduceOnly"] != "" {
		orderService.ReduceOnly(params["reduceOnly"] == "true")
	}
	if params["timeInForce"] != "" {
		orderService.TimeInForce(futures.TimeInForceType(strings.ToUpper(params["timeInForce"])))
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
	if params["selfTradePreventionMode"] != "" {
		orderService.SelfTradePreventionMode(futures.SelfTradePreventionMode(strings.ToUpper(params["selfTradePreventionMode"])))
	}
	if params["closePosition"] != "" {
		orderService.ClosePosition(params["closePosition"] == "true")
	}
	if params["activationPrice"] != "" {
		orderService.ActivationPrice(params["activationPrice"])
	}
	if params["callbackRate"] != "" {
		orderService.CallbackRate(params["callbackRate"])
	}
	if params["workingType"] != "" {
		orderService.WorkingType(futures.WorkingType(params["workingType"]))
	}
	if params["priceProtect"] != "" {
		orderService.PriceProtect(params["priceProtect"] == "true")
	}
	if params["newOrderRespType"] != "" {
		orderService.NewOrderResponseType(futures.NewOrderRespType(params["newOrderRespType"]))
	}
	if params["selfTradePreventionMode"] != "" {
		orderService.SelfTradePreventionMode(futures.SelfTradePreventionMode(params["selfTradePreventionMode"]))
	}

	return orderService.Do(context.Background())
}

func (c *Client) CancelOrder(symbol string, orderID int64, clientOrderID string) error {
	orderService := futures.NewClient(c.ApiKey, c.ApiSecret).NewCancelOrderService().Symbol(symbol)
	if orderID != 0 {
		orderService.OrderID(orderID)
	}
	if clientOrderID != "" {
		orderService.OrigClientOrderID(clientOrderID)
	}
	_, err := orderService.Do(context.Background())
	return err
}

func (c *Client) LeverageOrder(symbol string, leverage int) (*futures.SymbolLeverage, error) {
	orderService := futures.NewClient(c.ApiKey, c.ApiSecret).NewChangeLeverageService().Symbol(symbol)
	if leverage != 0 {
		orderService.Leverage(leverage)
	}
	return orderService.Do(context.Background())
}
