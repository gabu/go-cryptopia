package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetTradePairs(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetTradePairs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
