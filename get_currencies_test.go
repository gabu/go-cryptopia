package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetCurrencies(t *testing.T) {
	cryptopia := NewClient()
	ctx := context.Background()
	ret, err := cryptopia.GetCurrencies(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
