package cryptopia

import (
	"context"
	"fmt"
)

type MarketOrders map[string][]OpenOrder

type OpenOrder struct {
	TradePairID int
	Label       string
	Price       float64
	Volume      float64
	Total       float64
}

// GetMarketOrders returns the open buy and sell orders for the specified tarde pair.
func (c *Client) GetMarketOrders(ctx context.Context, market string, orderCount int) (MarketOrders, error) {
	path := "GetMarketOrders/" + market
	if orderCount != 0 {
		path += fmt.Sprintf("/%d", orderCount)
	}

	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return MarketOrders{}, err
	}

	var ret = &MarketOrders{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
