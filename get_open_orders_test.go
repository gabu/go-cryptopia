package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetOpenOrders(t *testing.T) {
	params := map[string]interface{}{
		"Market": "PAC_BTC",
	}
	cryptopia := newAuthClient()
	ctx := context.Background()
	ret, err := cryptopia.GetOpenOrders(ctx, params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
