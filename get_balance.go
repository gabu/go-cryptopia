package cryptopia

import (
	"context"

	"github.com/pkg/errors"
)

type Balances []Balance

type Balance struct {
	CurrencyID      int
	Symbol          string
	Total           float64
	Available       float64
	Unconfirmed     float64
	HeldForTrades   float64
	PendingWithdraw float64
	Address         string
	BaseAddress     string
	Status          string
	StatusMessage   string
}

// GetBalance returns all balances or a specific currency balance
func (c *Client) GetBalance(ctx context.Context) (Balances, error) {
	req, err := c.newAuthenticatedRequest(ctx, "GetBalance", nil)
	if err != nil {
		return Balances{}, errors.Wrap(err, "Faild to new authenticated request")
	}

	var ret = &Balances{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, errors.Wrap(err, "Faild to do request")
	}
	return *ret, nil
}
