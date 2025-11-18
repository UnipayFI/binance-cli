# Spot Module

## Quick Navigation
- [Account](#account--show-account-info)
- [Asset](#asset---show-account-assets)
- [Dividend](#dividend---query-asset-dividend-record)
- [Fee](#fee)
  - [Get BNB Burn Status](#fee---get-bnb-burn-status)
  - [Set BNB Burn Status](#fee---set-bnb-burn-status)
- [Order](#order)
  - [Create Market Order](#order---create-market-order)
  - [Create Limit Order](#order---create-limit-order)
  - [List open orders](#order---list-open-orders)
  - [List orders](#order---list-orders)
  - [Cancel order by ID](#order---cancel-order-by-id)
  - [Cancel all order by symbol](#order---cancel-all-order-by-symbol)

## Account — Show account info
Exec: `./binance-cli spot account`
```shell
┌────────────┬────────────────────┬───────────┬──────────────┬─────────────┬──────────────┬───────────────┬─────────────────────┐
│    UID     │  COMMISSION RATES  │ CAN TRADE │ CAN WITHDRAW │ CAN DEPOSIT │ ACCOUNT TYPE │  PERMISSIONS  │     UPDATE TIME     │
├────────────┼────────────────────┼───────────┼──────────────┼─────────────┼──────────────┼───────────────┼─────────────────────┤
│ xxxxxxxxxx │ Maker: 0.00000000  │ true      │ true         │ true        │ SPOT         │ [TRD_GRP_176] │ 2025-08-17 13:39:26 │
│            │ Taker: 0.00025000  │           │              │             │              │               │                     │
│            │ Buyer: 0.00000000  │           │              │             │              │               │                     │
│            │ Seller: 0.00000000 │           │              │             │              │               │                     │
└────────────┴────────────────────┴───────────┴──────────────┴─────────────┴──────────────┴───────────────┴─────────────────────┘
```

## Asset - Show account assets
Exec: `./binance-cli spot asset ls`
```shell
┌───────┬──────────────┬────────────┐
│ ASSET │    FREE      │   LOCKED   │
├───────┼──────────────┼────────────┤
│ BTC   │ 1.00000000   │ 0.00000000 │
│ USDT  │ 100.00000000 │ 0.00000000 │
└───────┴──────────────┴────────────┘
```

## Dividend - Query asset dividend record
Exec: `./binance-cli spot dividend ls`
```shell
┌─────────────────────┬───────┬──────────┬─────────────────────┬────────────────────┬────────────────┐
│         ID          │ ASSET │  AMOUNT  │    DIVIDEND TIME    │        INFO        │ TRANSACTION ID │
├─────────────────────┼───────┼──────────┼─────────────────────┼────────────────────┼────────────────┤
│ 4656120748007646527 │ USDT  │ 42.1329  │ 2025-08-18 02:40:22 │ BFUSD Daily reward │ 287113442430   │
│ 4654671248778998830 │ USDT  │ 108.6762 │ 2025-08-17 02:40:26 │ BFUSD Daily reward │ 286847043264   │
│ 4653237662432310020 │ USDT  │ 99.3702  │ 2025-08-16 02:56:17 │ BFUSD Daily reward │ 286622065819   │
│ 4651975355145888522 │ USDT  │ 116.2712 │ 2025-08-15 06:02:18 │ BFUSD Daily reward │ 286330628796   │
│ 4650581332527448370 │ USDT  │ 147.4873 │ 2025-08-14 06:57:28 │ BFUSD Daily reward │ 285921655320   │
└─────────────────────┴───────┴──────────┴─────────────────────┴────────────────────┴────────────────┘
```


## Fee
### Fee - Get BNB Burn Status
Exec: `./binance-cli spot fee status`
```shell
fee burn spotBNBBurn: true, interestBNBBurn: false
```

### Fee - Set BNB Burn Status
Exec: `./binance-cli spot fee set --spotBNBBurn=true --interestBNBBurn=true`
```shell
fee burn spotBNBBurn: true, interestBNBBurn: true
```

## Order
> support all docs parameters
> Docs Link: https://developers.binance.com/docs/binance-spot-api-docs/rest-api/trading-endpoints#new-order-trade

### Order - Create Market Order
Exec: `./binance-cli spot order create --symbol=BTCUSDT --side=BUY --type=MARKET --quantity=1.0`

### Order - Create Limit Order
Exec: `./binance-cli spot order create --symbol=SOLUSDT --side=BUY --type=LIMIT --quantity=0.2 --price=170 --timeInForce=GTC`

### Order - list open orders
Exec: `./binance-cli spot order open`
```shell
┌─────────────┬──────────────────────────────────┬─────────┬──────┬───────┬────────┬──────────────┬────────────┬───────────────────┐
│  ORDER ID   │         CLIENT ORDER ID          │ SYMBOL  │ SIDE │ TYPE  │ STATUS │    PRICE     │  QUANTITY  │ EXECUTED QUANTITY │
├─────────────┼──────────────────────────────────┼─────────┼──────┼───────┼────────┼──────────────┼────────────┼───────────────────┤
│ 13557621683 │ x-HNA2TXFJba82951e4dd3404daa2da4 │ SOLUSDT │ BUY  │ LIMIT │ NEW    │ 170.00000000 │ 0.20000000 │ 0.00000000        │
└─────────────┴──────────────────────────────────┴─────────┴──────┴───────┴────────┴──────────────┴────────────┴───────────────────┘
```

### Order - list orders
Exec: `./binance-cli spot order ls --symbol=SOLUSDT`
```shell
┌─────────────┬──────────────────────────────────┬─────────┬──────┬───────┬────────┬──────────────┬────────────┬───────────────────┐
│  ORDER ID   │         CLIENT ORDER ID          │ SYMBOL  │ SIDE │ TYPE  │ STATUS │    PRICE     │  QUANTITY  │ EXECUTED QUANTITY │
├─────────────┼──────────────────────────────────┼─────────┼──────┼───────┼────────┼──────────────┼────────────┼───────────────────┤
│ 13557621683 │ x-HNA2TXFJba82951e4dd3404daa2da4 │ SOLUSDT │ BUY  │ LIMIT │ NEW    │ 170.00000000 │ 0.20000000 │ 0.00000000        │
└─────────────┴──────────────────────────────────┴─────────┴──────┴───────┴────────┴──────────────┴────────────┴───────────────────┘
```

### Order - Cancel order by ID
Exec: `./binance-cli spot order cancel --origClientOrderId=xxxxx --symbol=SOLUSDT`
OR
Exec: `./binance-cli spot order cancel --clientId=xxxxx --symbol=SOLUSDT`

### Order - Cancel all order by symbol
Exec: `./binance-cli spot order cancel --symbol=SOLUSDT`