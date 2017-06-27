package cryptopia

import (
	"context"
)

type TradePairs []TradePair

type TradePair struct {
	ID               int
	Label            string
	Currency         string
	Symbol           string
	BaseCurrency     string
	BaseSymbol       string
	Status           string
	StatusMessage    string
	TradeFee         float64
	MinimumTrade     float64
	MaximumTrade     float64
	MinimumBaseTrade float64
	MaximumBaseTrade float64
	MinimumPrice     float64
	MaximumPrice     float64
}

// GetTradePairs returns all trade pair data.
func (c *Client) GetTradePairs(ctx context.Context) (TradePairs, error) {
	req, err := c.newRequest(ctx, "GET", "GetTradePairs", nil)
	if err != nil {
		return TradePairs{}, err
	}

	var ret = &TradePairs{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
