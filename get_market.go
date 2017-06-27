package cryptopia

import (
	"context"
	"fmt"
)

type Market struct {
	TradePairID    int
	Label          string
	AskPrice       float64
	BidPrice       float64
	Low            float64
	High           float64
	Volume         float64
	LastPrice      float64
	BuyVolume      float64
	SellVolume     float64
	Change         float64
	Open           float64
	Close          float64
	BaseVolume     float64
	BuyBaseVolume  float64
	SellBaseVolume float64
}

// GetMarket returns market data for the specified trade pair.
func (c *Client) GetMarket(ctx context.Context, market string, hours int) (Market, error) {
	path := "GetMarket/" + market
	if hours != 0 {
		path += fmt.Sprintf("/%d", hours)
	}

	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return Market{}, err
	}

	var ret = &Market{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
