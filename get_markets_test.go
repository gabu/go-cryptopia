package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarkets(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetMarkets(ctx, "BTC", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
