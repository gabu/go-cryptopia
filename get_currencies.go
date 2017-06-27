package cryptopia

import (
	"context"
)

type Currencies []Currency

type Currency struct {
	ID                   int
	Name                 string
	Symbol               string
	Algorithm            string
	WithdrawFee          float64
	MinWithdraw          float64
	MinBaseTrade         float64
	IsTipEnabled         bool
	MinTip               float64
	DepositConfirmations int
	Status               string
	StatusMessage        string
	ListingStatus        string
}

// GetCurrencies returns all currency data.
func (c *Client) GetCurrencies(ctx context.Context) (Currencies, error) {
	req, err := c.newRequest(ctx, "GET", "GetCurrencies", nil)
	if err != nil {
		return Currencies{}, err
	}

	var ret = &Currencies{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
