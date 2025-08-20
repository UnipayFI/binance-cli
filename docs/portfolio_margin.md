# Portfolio Margin Module

## Quick Navigation
- [Interest](#interest---get-margin-borrowloan-interest-history)
- [Loan](#loan)
  - [Loan exec](#loan---loan-exec)
  - [Loan list](#loan---loan-list)
  - [Repay for a margin loan](#loan---repay-for-a-margin-loan)
  - [Repay debt for a margin loan](#loan---repay-debt-for-a-margin-loan)
  - [Repay history](#loan---repay-history)
- [Order](#order)
  - [Create Market Order](#order---create-market-order)
  - [Create Limit Order](#order---create-limit-order)
  - [Cancel order by ID](#order---cancel-order-by-id)
  - [Cancel all order by symbol](#order---cancel-all-order-by-symbol)
  - [Query All Margin Account Orders](#order---query-all-margin-account-orders)
  - [Query Current Margin Open Order](#order---query-current-margin-open-order)

## Interest - Get Margin Borrow/Loan Interest History
Exec: `./binance-cli portfolio margin interest-history`
```shell
┌─────────────────────┬───────┬───────────┬───────────┬───────────┬────────────┬───────────────┬───────────────────────┐
│        TX ID        │ ASSET │ RAW ASSET │   TYPE    │ PRINCIPAL │  INTEREST  │ INTEREST RATE │ INTEREST ACCURED TIME │
├─────────────────────┼───────┼───────────┼───────────┼───────────┼────────────┼───────────────┼───────────────────────┤
│ 1352286576452864727 │ USDT  │ USDT      │ ON_BORROW │ 45.3313   │ 0.00024995 │ 0.00013233    │ 2022-12-28 01:00:00   │
└─────────────────────┴───────┴───────────┴───────────┴───────────┴────────────┴───────────────┴───────────────────────┘
```

## Loan

### Loan - Loan exec
Exec: `./binance-cli portfolio margin loan exec --amount=1.0 --asset=SOL`

### Loan - Loan list
Exec: `./binance-cli portfolio margin loan ls --asset=SOL`
```shell
┌──────────────┬───────┬───────────┬───────────┬─────────────────────┐
│    TX ID     │ ASSET │ PRINCIPAL │  STATUS   │      TIMESTAMP      │
├──────────────┼───────┼───────────┼───────────┼─────────────────────┤
│ 287807735989 │ SOL   │ 1.0       │ CONFIRMED │ 2025-08-20 14:07:27 │
└──────────────┴───────┴───────────┴───────────┴─────────────────────┘
```

### Loan - Repay for a margin loan.
Exec: `./binance-cli portfolio margin loan repay --amount=1.0 --asset=SOL`

### Loan - Repay debt for a margin loan.
Exec: `./binance-cli portfolio margin loan repay-debt --asset=SOL --amount=0.01 --specifyRepayAssets=USDT`

### Loan - Repay history
Exec: `./binance-cli portfolio margin loan repay-list --asset=SOL`
```shell
┌──────────────┬───────┬────────┬─────────────┬────────────┬───────────┬─────────────────────┐
│    TX ID     │ ASSET │ AMOUNT │  INTEREST   │ PRINCIPAL  │  STATUS   │      TIMESTAMP      │
├──────────────┼───────┼────────┼─────────────┼────────────┼───────────┼─────────────────────┤
│ 287810116588 │ SOL   │ 0.01   │ 0.000000080 │ 0.00999992 │ CONFIRMED │ 2025-08-20 14:19:07 │
│ 287809473014 │ SOL   │ 0.01   │ 0.000000080 │ 0.00999992 │ CONFIRMED │ 2025-08-20 14:16:11 │
│ 287808117953 │ SOL   │ 0.01   │ 0.000000080 │ 0.00999992 │ CONFIRMED │ 2025-08-20 14:09:15 │
└──────────────┴───────┴────────┴─────────────┴────────────┴───────────┴─────────────────────┘
```

## Order

### Order - Create Market Order
Exec: `./binance-cli portfolio margin order create --symbol=SOLUSDT  --quantity=1.0 --side=SELL --type=MARKET`

### Order - Create Limit Order
Exec: `./binance-cli portfolio margin order create --symbol=SOLSUDT  --quantity=1.0 --side=SELL --type=LIMIT --price=200.0 --timeInForce=GTC`

### Order - Cancel order by ID
Exec: `./binance-cli portfolio margin order cancel --symbol=SOLSUDT --orderId=xxx`
OR
Exec: `./binance-cli portfolio margin order cancel --symbol=SOLSUDT --clientOrderId=xxx`

### order - Cancel all order by symbol
Exec: `./binance-cli portfolio margin order cancel --symbol=SOLSUDT`

### Order - Query All Margin Account Orders
Exec: `./binance-cli portfolio margin order ls --symbol=BFUSDUSDT`
```shell
┌──────────┬───────────┬──────┬──────────┬────────────┬─────────────┬───────────────────┬─────────────────────┬─────────────────────┐
│ ORDER ID │  SYMBOL   │ SIDE │  STATUS  │   PRICE    │  QUANTITY   │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├──────────┼───────────┼──────┼──────────┼────────────┼─────────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 69259    │ BFUSDUSDT │ SELL │ FILLED   │ 0.00000000 │ 20.00000000 │ 20.00000000       │ 2025-08-20 14:29:28 │ 2025-08-20 14:29:28 │
│ 69268    │ BFUSDUSDT │ SELL │ CANCELED │ 1.00000000 │ 20.00000000 │ 0.00000000        │ 2025-08-20 14:30:49 │ 2025-08-20 14:34:47 │
│ 69335    │ BFUSDUSDT │ SELL │ CANCELED │ 1.00000000 │ 20.00000000 │ 0.00000000        │ 2025-08-20 14:40:10 │ 2025-08-20 14:40:12 │
└──────────┴───────────┴──────┴──────────┴────────────┴─────────────┴───────────────────┴─────────────────────┴─────────────────────┘
```

### Order - Query Current Margin Open Order
Exec: `./binance-cli portfolio margin order open --symbol=BFUSDUSDT`
```shell
┌──────────┬───────────┬──────┬────────┬────────────┬─────────────┬───────────────────┬─────────────────────┬─────────────────────┐
│ ORDER ID │  SYMBOL   │ SIDE │ STATUS │   PRICE    │  QUANTITY   │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├──────────┼───────────┼──────┼────────┼────────────┼─────────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 69348    │ BFUSDUSDT │ SELL │ NEW    │ 1.00000000 │ 20.00000000 │ 0.00000000        │ 2025-08-20 14:43:33 │ 2025-08-20 14:43:33 │
└──────────┴───────────┴──────┴────────┴────────────┴─────────────┴───────────────────┴─────────────────────┴─────────────────────┘
```