package cryptopia

import (
	"context"
	"fmt"
)

type MarketOrderGroups []MarketOrderGroup

type MarketOrderGroup struct {
	TradePairID int
	Market      string
	Buy         []OpenOrder
	Sell        []OpenOrder
}

// GetMarketOrderGroups returns the open buy and sell orders for the specified markets.
func (c *Client) GetMarketOrderGroups(ctx context.Context, markets string, orderCount int) (MarketOrderGroups, error) {
	path := "GetMarketOrderGroups/" + markets
	if orderCount != 0 {
		path += fmt.Sprintf("/%d", orderCount)
	}

	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return MarketOrderGroups{}, err
	}

	var ret = &MarketOrderGroups{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
