package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarketOrderGroups(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetMarketOrderGroups(ctx, "DOT_BTC-LTC-DOT", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
