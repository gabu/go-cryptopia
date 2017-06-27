# go-cryptopia

An unofficial [Cryptopia Public and Private API](https://www.cryptopia.co.nz/Forum/Category/45) client for Go.

## Supports

### Public API

- [x] GetCurrencies
- [x] GetTradePairs
- [x] GetMarkets
- [x] GetMarket
- [x] GetMarketHistory
- [x] GetMarketOrders
- [x] GetMarketOrderGroups

### Private API (needs an authentication)

- [x] GetBalance
- [ ] GetDepositAddress
- [ ] GetOpenOrders
- [ ] GetTradeHistory
- [ ] GetTransactions
- [ ] SubmitTrade
- [ ] CancelTrade
- [ ] SubmitTip
- [ ] SubmitWithdraw
- [ ] SubmitTransfer

## Usage

### GetCurrencies

```go
package main

import (
	"context"
	"fmt"

	"github.com/gabu/go-cryptopia"
)

func main() {
	cryptopia := cryptopia.NewClient()
	ctx := context.Background()
	currencies, err := cryptopia.GetCurrencies(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(currencies)
}
```

### Authentication

```go
func main() {
	cryptopia := cryptopia.NewClient().Auth("YOUR API KEY", "YOUR API SECRET")
}
```
