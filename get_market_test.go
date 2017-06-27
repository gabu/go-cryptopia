package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetMarket(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetMarket(ctx, "SKY_BTC", 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
