package cryptopia

import (
	"context"
	"fmt"
)

type Markets []Market

// GetMarkets returns all market data.
func (c *Client) GetMarkets(ctx context.Context, baseMarket string, hours int) (Markets, error) {
	path := "GetMarkets"
	if baseMarket != "" {
		path += "/" + baseMarket
	}
	if hours != 0 {
		path += fmt.Sprintf("/%d", hours)
	}

	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return Markets{}, err
	}

	var ret = &Markets{}
	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}
	return *ret, nil
}
