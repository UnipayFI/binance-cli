package unified

import (
	"context"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2/futures"
	"github.com/adshao/go-binance/v2/portfolio"
)

func (c *Client) GetOrders(symbol string) (OrderList, error) {
	orders, err := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewUMAllOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}
	orderList := make(OrderList, len(orders))
	for i, order := range orders {
		orderList[i] = *order
	}
	return orderList, nil
}

func (c *Client) CreateUMOrder(params map[string]string) (*portfolio.UMOrder, error) {
	sideType := portfolio.SideType(strings.ToUpper(params["side"]))
	t := portfolio.OrderType(strings.ToUpper(params["type"]))
	orderService := portfolio.NewClient(c.ApiKey, c.ApiSecret).NewUMOrderService().Symbol(params["symbol"]).Side(sideType).Type(t)

	if params["positionSide"] != "" {
		orderService.PositionSide(portfolio.PositionSideType(params["positionSide"]))
	}
	if params["timeInForce"] != "" {
		orderService.TimeInForce(portfolio.TimeInForceType(params["timeInForce"]))
	}
	if params["quantity"] != "" {
		orderService.Quantity(params["quantity"])
	}
	if params["reduceOnly"] != "" {
		orderService.ReduceOnly(params["reduceOnly"] == "true")
	}
	if params["price"] != "" {
		orderService.Price(params["price"])
	}
	if params["newClientOrderId"] != "" {
		orderService.NewClientOrderID(params["newClientOrderId"])
	}
	if params["newOrderRespType"] != "" {
		orderService.NewOrderRespType(portfolio.NewOrderRespType(params["newOrderRespType"]))
	}
	if params["priceMatch"] != "" {
		orderService.PriceMatch(portfolio.PriceMatchType(params["priceMatch"]))
	}
	if params["selfTradePreventionMode"] != "" {
		orderService.SelfTradePreventionMode(portfolio.SelfTradePreventionMode(params["selfTradePreventionMode"]))
	}
	if params["goodTillDate"] != "" {
		goodTillDate, err := strconv.ParseInt(params["goodTillDate"], 10, 64)
		if err != nil {
			return nil, err
		}
		orderService.GoodTillDate(goodTillDate)
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
