package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type GetOpenOrders []GetOpenOrder

type GetOpenOrder struct {
	OrderID     int
	TradePairId int
	Market      string
	Type        string
	Rate        float64
	Amount      float64
	Total       float64
	Remaining   float64
	TimeStamp   string
}

// GetOpenOrders returns all open orders
func (c *Client) GetOpenOrders(ctx context.Context, market map[string]interface{}) (GetOpenOrders, error) {
	req, err := c.newAuthenticatedRequest(ctx, "GetOpenOrders", market)
	if err != nil {
		return GetOpenOrders{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &GetOpenOrders{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Faild to do request")
	}
	return *ret, nil
}
