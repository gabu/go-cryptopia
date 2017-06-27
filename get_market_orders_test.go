package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarketOrders(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetMarketOrders(ctx, "SKY_BTC", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
