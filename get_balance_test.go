package cryptopia

import (
	"context"
	"fmt"
	"testing"
)

func TestGetBalance(t *testing.T) {
	cryptopia := newAuthClient()
	ctx := context.Background()
	ret, err := cryptopia.GetBalance(ctx)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
