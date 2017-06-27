package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarketHistory(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetMarketHistory(ctx, "SKY_BTC", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
