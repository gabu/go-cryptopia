package cryptopia

import (
	"context"
	"fmt"
)

type MarketHistory []ClosedOrder

type ClosedOrder struct {
	TradePairID int
	Label       string
	Type        string
	Price       float64
	Amout       float64
	Total       float64
	Timestamp   int
}

// GetMarketHistory returns the market history data for the specified tarde pair.
func (c *Client) GetMarketHistory(ctx context.Context, market string, hours int) (MarketHistory, error) {
	path := "GetMarketHistory/" + market
	if hours != 0 {
		path += fmt.Sprintf("/%d", hours)
	}

	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return MarketHistory{}, err
	}

	var ret = &MarketHistory{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
