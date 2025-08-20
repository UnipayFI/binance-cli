# Portfolio USDⓈ-Margined Module

## Quick Navigation
- [Commission-rate](#commission-rate---get-user-commission-rate-for-um)
- [Fee](#fee)
  - [Show BNB burn status](#fee---show-bnb-burn-status)
  - [Set BNB burn status](#fee---set-bnb-burn-status)
- [Income](#income---get-um-income-history)
- [Order](#order)
  - [Create Market Order](#order---create-market-order)
  - [Create Limit Order](#order---create-limit-order)
  - [Reduce short positions(Market)](#order---reduce-short-positionsmarket)
  - [Reduce long positions(Limit)](#order---reduce-long-positionslimit)
  - [List open orders](#order---list-open-orders)
  - [List orders](#order---list-orders)
  - [Cancel order by ID](#order---cancel-order-by-id)
  - [Cancel all order by symbol](#order---cancel-all-order-by-symbol)
  - [Download Order](#order---download-order)
- [Position](#position)
  - [List positions](#position---list-positions)
  - [Show position risk](#position---show-position-risk)
  - [Set position side](#position---set-position-side)
  - [Show position side](#position---show-position-side)
- [Symbol](#symbol)
  - [Set leverage](#symbol---set-leverage)
  - [Show symbol config](#symbol---show-symbol-config)

## Commission-rate - Get User Commission Rate for UM
Exec: `./binance-cli portfolio um commission-rate --symbol=BTCUSDT`
```shell
BTCUSDT commission mraker rate: -0.000030, taker rate: 0.000200
```

## Fee
### Fee - Show BNB burn status
Exec: `./binance-cli portfolio um fee status`
```shell
fee burn status: true
```

### Fee - Set BNB burn status
Exec: `./binance-cli portfolio um fee set --status=true`
```shell
fee burn changed to: true
```

## Income - Get UM Income History
Exec: `./binance-cli portfolio um income`
```shell
┌────────────────────┬────────────┬─────────┬──────────────┬──────────────┬───────┬─────────────┬─────────────────────┐
│      TRAN ID       │  TRADE ID  │ SYMBOL  │ INCOME TYPE  │    INCOME    │ ASSET │    INFO     │        TIME         │
├────────────────────┼────────────┼─────────┼──────────────┼──────────────┼───────┼─────────────┼─────────────────────┤
│ 900616219512893    │ 6219512893 │ ETHUSDT │ COMMISSION   │ 0.15003511   │ USDT  │ 6219512893  │ 2025-08-19 05:34:27 │
│ 900616219513204    │ 6219513204 │ ETHUSDT │ REALIZED_PNL │ 19.96812556  │ USDT  │ 6219513204  │ 2025-08-19 05:34:35 │
│ 100957777041736630 │            │ BTCUSDT │ FUNDING_FEE  │ 12.76719674  │ USDT  │ FUNDING_FEE │ 2025-08-19 08:00:00 │
│ 100957777964483500 │            │ ETHUSDT │ FUNDING_FEE  │ 11.34082844  │ USDT  │ FUNDING_FEE │ 2025-08-19 08:00:00 │
│ 100957782452388400 │            │ SOLUSDT │ FUNDING_FEE  │ 34.06565647  │ USDT  │ FUNDING_FEE │ 2025-08-19 08:00:00 │
└────────────────────┴────────────┴─────────┴──────────────┴──────────────┴───────┴─────────────┴─────────────────────┘
```

## Order

### Order - Create Market Order
Exec: `./binance-cli portfolio um order create --symbol=SOLUSDT --side=SELL --positionSide=SHORT --type=MARKET --quantity=0.1`

### Order - Create Limit Order
Exec: `./binance-cli portfolio um order create --symbol=SOLUSDT --side=BUY --positionSide=LONG --type=LIMIT --timeInForce=GTC --quantity=0.1 --price=180`

### Order - Reduce short positions(Market)
Exec: `./binance-cli portfolio um order create --symbol=SOLUSDT --side=BUY --positionSide=SHORT --type=MARKET --quantity=0.1`

### Order - Reduce long positions(Limit)
Exec: `./binance-cli portfolio um order create --symbol=SOLUSDT --side=SELL --positionSide=LONG --type=LIMIT --timeInForce=GTC --price=4000.0 --quantity=0.1`

### Order - list open orders
Exec: `./binance-cli portfolio um order open`
```shell
┌──────────────┬─────────┬──────┬────────┬──────────┬──────────┬───────────────────┬─────────────────────┬─────────────────────┐
│   ORDER ID   │ SYMBOL  │ SIDE │ STATUS │  PRICE   │ QUANTITY │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├──────────────┼─────────┼──────┼────────┼──────────┼──────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 137907377496 │ SOLUSDT │ SELL │ NEW    │ 182.0000 │ 0.10     │ 0.10              │ 2025-08-20 15:25:01 │ 2025-08-20 15:25:01 │
└──────────────┴─────────┴──────┴────────┴──────────┴──────────┴───────────────────┴─────────────────────┴─────────────────────┘
```

### Order - list orders
Exec: `./binance portfolio um order ls`
```shell
┌──────────────┬─────────┬──────┬──────────┬──────────┬──────────┬───────────────────┬─────────────────────┬─────────────────────┐
│   ORDER ID   │ SYMBOL  │ SIDE │  STATUS  │  PRICE   │ QUANTITY │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├──────────────┼─────────┼──────┼──────────┼──────────┼──────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 137425039598 │ SOLUSDT │ BUY  │ FILLED   │ 0.0000   │ 0.70     │ 0.70              │ 2025-08-18 16:02:31 │ 2025-08-18 16:02:31 │
│ 137904175369 │ SOLUSDT │ SELL │ FILLED   │ 0.0000   │ 0.10     │ 0.10              │ 2025-08-20 15:05:32 │ 2025-08-20 15:05:32 │
│ 137904369902 │ SOLUSDT │ BUY  │ FILLED   │ 0.0000   │ 0.10     │ 0.10              │ 2025-08-20 15:06:28 │ 2025-08-20 15:06:28 │
│ 137906762985 │ SOLUSDT │ SELL │ FILLED   │ 0.0000   │ 0.10     │ 0.10              │ 2025-08-20 15:20:45 │ 2025-08-20 15:20:45 │
└──────────────┴─────────┴──────┴──────────┴──────────┴──────────┴───────────────────┴─────────────────────┴─────────────────────┘
```

### Order - Cancel order by ID
Exec: `./binance portfolio um order cancel --symbol=SOLUSDT --orderId=xxx`
OR
`./binance portfolio um order cancel --symbol=SOLUSDT --origClientOrderId=xxx`

### Order - Cancel all order by symbol
Exec: `./binance portfolio um order cancel --symbol=SOLUSDT`

### order - Download Order
Exec: `./binance-cli portfolio um order download --startTime=1755588612 --endTime=1755675012`
```shell
downloadID: 1010450789015306240
waiting for processing...
waiting for processing...
waiting for processing...
waiting for processing...
waiting for processing...
waiting for processing...
waiting for processing...
download link: https://bin-prod-user-rebate-bucket.s3.amazonaws.com/data-download-task/xxxxxx
```

## Position

### Position - List positions
Exec: `./binance-cli portfolio um position ls`
```shell
┌─────────┬────────────────┬────────────────────┬───────────────────┬───────────────┬─────────────────┬──────────────────┬─────────────────────┐
│ SYMBOL  │ INITIAL MARGIN │ MAINTENANCE MARGIN │ UNREALIZED PROFIT │ POSITION SIDE │ POSITION AMOUNT │     NOTIONAL     │     UPDATE TIME     │
├─────────┼────────────────┼────────────────────┼───────────────────┼───────────────┼─────────────────┼──────────────────┼─────────────────────┤
│ SOLUSDT │ 83561.57310000 │ 6881.15731000      │ 7418.45390653     │ SHORT         │ -4625.35        │ -835615.73100000 │ 2025-08-19 23:40:26 │
└─────────┴────────────────┴────────────────────┴───────────────────┴───────────────┴─────────────────┴──────────────────┴─────────────────────┘
```

### Position - Show position risk
Exec: `./binance-cli portfolio um position risk`
```shell
┌─────────┬────────────────┬────────────────┬───────────────────┬───────────────────┬──────────┬──────────────┬───────────────┬──────────┬──────────────────┬─────────────────────┐
│ SYMBOL  │ POSITION AMONT │  ENTRY PRICE   │ UNREALIZED PROFIT │ LIQUIDATION PRICE │ LEVERAGE │ MAX NOTIONAL │ POSITION SIDE │ NOTIONAL │   UPDATE TIME    │                     │
├─────────┼────────────────┼────────────────┼───────────────────┼───────────────────┼──────────┼──────────────┼───────────────┼──────────┼──────────────────┼─────────────────────┤
│ SOLUSDT │ -4625.35       │ 182.2638686629 │ 7502.90712176     │ 307.4485281       │ 10       │              │ 40000000      │ SHORT    │ -835531.27779818 │ 2025-08-19 23:40:26 │
└─────────┴────────────────┴────────────────┴───────────────────┴───────────────────┴──────────┴──────────────┴───────────────┴──────────┴──────────────────┴─────────────────────┘
```

### Position - Set position side
Exec: `./binance-cli portfolio um position set-side --dualSidePosition=true`

### Position - Show position side
Exec: `./binance-cli portfolio um position side`
```shell
dual side position: true
```

## Symbol

### Symbol - set-leverage
Exec: `./binance-cli portfolio um symbol set-leverage --leverage=10 --symbol=BTCUSDT`

### Symbol - Show symbol config
Exec: `./binance-cli portfolio um symbol show --symbol=BTCUSDT`
```shell
┌─────────────────┬─────────────┬────────────────────┬──────────┬────────────────────┐
│     SYMBOL      │ MARGIN TYPE │ IS AUTO ADD MARGIN │ LEVERAGE │ MAX NOTIONAL VALUE │
├─────────────────┼─────────────┼────────────────────┼──────────┼────────────────────┤
│ BTCUSDT         │ CROSSED     │ false              │ 10       │ 100000000          │
└─────────────────┴─────────────┴────────────────────┴──────────┴────────────────────┘
```

