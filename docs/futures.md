# Futures Module

## Quick Navigation
- [Account](#account)
  - [Show account config](#account---show-account-config)
  - [Show account balances](#account---show-account-balances)
- [Commission-rate](#commission-rate---show-commission-rate)
- [Fee](#fee)
  - [Get BNB Burn Status](#fee---get-bnb-burn-status)
  - [Change user's BNB Fee Discount](#fee---change-users-bnb-fee-discount-on-every-symbol)
- [Income](#income---query-income-history)
- [Multi-assets-mode](#multi-assets-mode)
  - [Show multi-assets mode](#multi-assets-mode---show-multi-assets-mode)
  - [Set multi-assets mode](#multi-assets-mode---set-multi-assets-mode)
- [Order](#order)
  - [Create Market Order](#order---create-market-order)
  - [Create Limit Order](#order---create-limit-order)
  - [Reduce short positions(Market)](#order---reduce-short-positionsmarket)
  - [Reduce long positions(Limit)](#order---reduce-long-positionslimit)
  - [List open orders](#order---list-open-orders)
  - [List orders](#order---list-orders)
  - [Cancel order by ID](#order---cancel-order-by-id)
  - [Cancel all order by symbol](#order---cancel-all-order-by-symbol)
- [Position](#position)
  - [List positions](#position---list-positions)
  - [Show position risk](#position---show-position-risk)
  - [Set position margin](#position---set-position-margin)
  - [Change Position Mode](#position---change-position-mode)
  - [Get user's position mode](#position---get-users-position-mode-on-every-symbol)
- [Symbol](#symbol)
  - [Change Initial Leverage](#symbol---change-initial-leverage)
  - [Set margin type](#symbol---set-margin-type)
  - [Show symbol config](#symbol---show-symbol-config)
- [Trade](#trade----get-trades-for-a-specific-account-and-symbol)

## Account - Show account config
Exec: ` ./binance-cli futures account config`
```shell
┌──────────┬───────────┬─────────────┬──────────────┬────────────────────┬─────────────────────┬────────────────┐
│ FEE TIER │ CAN TRADE │ CAN DEPOSIT │ CAN WITHDRAW │ DUAL SIDE POSITION │ MULTI ASSETS MARGIN │ TRADE GROUP ID │
├──────────┼───────────┼─────────────┼──────────────┼────────────────────┼─────────────────────┼────────────────┤
│ 8        │ true      │ true        │ true         │ true               │ false               │ -1             │
└──────────┴───────────┴─────────────┴──────────────┴────────────────────┴─────────────────────┴────────────────┘
```

## Account - Show account balances
Exec: `./binance-cli futures account balances`
```shell
┌───────┬────────────────┬──────────────────────┬────────────────┬───────────────────┬─────────────────────┐
│ ASSET │    BALANCE     │ CROSS WALLET BALANCE │  CROSS UN PNL  │ AVAILABLE BALANCE │ MAX WITHDRAW AMOUNT │
├───────┼────────────────┼──────────────────────┼────────────────┼───────────────────┼─────────────────────┤
│ BNB   │ 0.09949690     │ 0.09949690           │ 0.00000000     │ 0.09949690        │ 0.09949690          │
│ USDC  │ 0.70000000     │ 0.70000000           │ 0.00000000     │ 0.70000000        │ 0.70000000          │
└───────┴────────────────┴──────────────────────┴────────────────┴───────────────────┴─────────────────────┘
```

## Commission-rate - Show commission rate
Exec: `./binance-cli futures commission-rate --symbol=BTCUSDT`
```shell
┌─────────┬───────────────────────┬───────────────────────┐
│ SYMBOL  │ MAKER COMMISSION RATE │ TAKER COMMISSION RATE │
├─────────┼───────────────────────┼───────────────────────┤
│ BTCUSDT │ -0.0030%              │ 0.0200%               │
└─────────┴───────────────────────┴───────────────────────┘
```

## Fee - Get BNB Burn Status
Exec: `./binance-cli futures fee status`

## Fee - Change user's BNB Fee Discount on EVERY symbol
Exec: `./binance-cli futures fee set --feeBurn=true`

## Income - Query income history
Exec: `/binance-cli futures income`
```shell
┌───────┬─────────────────┬──────────────┬─────────────┬─────────┬─────────────────────┬───────────────────┬────────────┐
│ ASSET │     INCOME      │ INCOME TYPE  │    INFO     │ SYMBOL  │        TIME         │      TRAN ID      │  TRADE ID  │
├───────┼─────────────────┼──────────────┼─────────────┼─────────┼─────────────────────┼───────────────────┼────────────┤
│ USDT  │ 0.02824612      │ COMMISSION   │ 6560390403  │ BTCUSDT │ 2025-08-16 13:52:58 │ 900316560390403   │ 6560390403 │
│ USDT  │ -3.81817128     │ REALIZED_PNL │ 6560390646  │ BTCUSDT │ 2025-08-16 13:53:00 │ 900316560390646   │ 6560390646 │
│ USDT  │ 0.02824173      │ COMMISSION   │ 6560390646  │ BTCUSDT │ 2025-08-16 13:53:00 │ 900316560390646   │ 6560390646 │
└───────┴─────────────────┴──────────────┴─────────────┴─────────┴─────────────────────┴───────────────────┴────────────┘
```

## Multi-assets-mode - Show multi-assets mode
Exec: `./binance-cli futures multi-assets-mode show`

## Multi-assets-mode - Set multi-assets mode
Exec: `./binance-cli futures multi-assets-mode set --multiAssetsMargin=true`

## Order
> support all docs parameters
> Docs Link: https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/New-Order-Trade

### Order - Create Market Order
Exec: `./binance-cli futures order create --symbol=ETHUSDT --side=SELL --positionSide=SHORT --type=MARKET --quantity=0.1`

### Order - Create Limit Order
Exec: `./binance-cli futures order create --symbol=ETHUSDT --side=BUY --positionSide=LONG --type=LIMIT --timeInForce=GTC --quantity=0.01 --price=4000`

### Order - Reduce short positions(Market)
Exec: `./binance-cli futures order create --symbol=ETHUSDT --side=BUY --positionSide=SHORT --type=MARKET --quantity=1.0`

### Order - Reduce long positions(Limit)
Exec: `./binance-cli futures order create --symbol=ETHUSDT --side=SELL --positionSide=LONG --type=LIMIT --timeInForce=GTC --price=4000.0 --quantity=0.01`

### Order - list open orders
Exec: `./binance-cli futures order open`
```shell
┌─────────────────────┬─────────┬──────┬───────────────┬────────┬───────┬──────────┬───────────────────┬─────────────────────┬─────────────────────┐
│      ORDER ID       │ SYMBOL  │ SIDE │ POSITION SIDE │ STATUS │ PRICE │ QUANTITY │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├─────────────────────┼─────────┼──────┼───────────────┼────────┼───────┼──────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 8389765940941702977 │ ETHUSDT │ BUY  │ LONG          │ NEW    │ 4000  │ 0.010    │ 0                 │ 2025-08-19 15:32:41 │ 2025-08-19 15:32:41 │
└─────────────────────┴─────────┴──────┴───────────────┴────────┴───────┴──────────┴───────────────────┴─────────────────────┴─────────────────────┘
```

### Order - list orders
Exec: `./binance-cli futures order ls`
```shell
┌─────────────────────┬─────────┬──────┬───────────────┬──────────┬───────┬──────────┬───────────────────┬─────────────────────┬─────────────────────┐
│      ORDER ID       │ SYMBOL  │ SIDE │ POSITION SIDE │  STATUS  │ PRICE │ QUANTITY │ EXECUTED QUANTITY │        TIME         │     UPDATE TIME     │
├─────────────────────┼─────────┼──────┼───────────────┼──────────┼───────┼──────────┼───────────────────┼─────────────────────┼─────────────────────┤
│ 8389765940941702977 │ ETHUSDT │ BUY  │ LONG          │ CANCELED │ 4000  │ 0.010    │ 0                 │ 2025-08-19 15:32:41 │ 2025-08-19 15:33:17 │
└─────────────────────┴─────────┴──────┴───────────────┴──────────┴───────┴──────────┴───────────────────┴─────────────────────┴─────────────────────┘
```

### Order - Cancel order by ID
Exec: `./binance-cli futures order cancel --symbol=ETHUSDT --orderId=xxx`
OR
`./binance-cli futures order cancel --symbol=ETHUSDT --origClientOrderId=xxx`

### Order - Cancel all order by symbol
Exec: `./binance-cli futures order cancel --symbol=ETHUSDT`

## Position

### Position - List positions
Exec: `./binance-cli futures position ls`
```shell
┌─────────┬───────────────┬─────────────────┬────────────────┬───────────────────┬──────────┬─────────────────────┐
│ SYMBOL  │ POSITION SIDE │ POSITION AMOUNT │  ENTRY PRICE   │ UNREALIZED PROFIT │ LEVERAGE │     UPDATE TIME     │
├─────────┼───────────────┼─────────────────┼────────────────┼───────────────────┼──────────┼─────────────────────┤
│ BTCUSDT │ SHORT         │ -0.186          │ 117196.6285896 │ 443.67111766      │ 10       │ 2025-08-18 15:53:21 │
└─────────┴───────────────┴─────────────────┴────────────────┴───────────────────┴──────────┴─────────────────────┘
```

### Position - Show position risk
Exec: `./binance-cli futures position risk`
```shell
┌─────────┬───────────────┬─────────────────┬─────────────────┬────────────────┬─────────────────┬───────────────────┬───────────────────┬─────────────────────┐
│ SYMBOL  │ POSITION SIDE │ POSITION AMOUNT │    NOTIONAL     │  ENTRY PRICE   │   MARK PRICE    │ UNREALIZED PROFIT │ LIQUIDATION PRICE │     UPDATE TIME     │
├─────────┼───────────────┼─────────────────┼─────────────────┼────────────────┼─────────────────┼───────────────────┼───────────────────┼─────────────────────┤
│ BTCUSDT │ SHORT         │ -0.186 BTC      │ -21345.25499086 │ 117196.6285896 │ 114759.43543478 │ 453.31792679      │ 536665.62249500   │ 2025-08-18 15:53:21 │
└─────────┴───────────────┴─────────────────┴─────────────────┴────────────────┴─────────────────┴───────────────────┴───────────────────┴─────────────────────┘
```

### Position - Set position margin
Exec: `./binance-cli futures position set-margin --amount=1.0 --positionSide=SHORT --symbol=BTCUSDT --type=ADD`

### Position - Change Position Mode
Exec: `./binance-cli futures position set-side --dualSidePosition=true`

### Position - Get user's position mode on EVERY symbol
Exec: `./binance-cli futures position side`

## Symbol

### Symbol - Change Initial Leverage
Exec: `./binance-cli futures symbol set-leverage --symbol=BTCUSDT --leverage=10`

### Symbol - Set margin type
Exec: `./binance-cli futures symbol set-margin-type --symbol=BTCUSDT --marginType=CROSSED`

### Symbol - Show symbol config
Exec: ` ./binance-cli futures symbol show --symbol=BTCUSDT`
```shell
┌─────────┬─────────────┬────────────────────┬──────────┬────────────────────┐
│ SYMBOL  │ MARGIN TYPE │ IS AUTO ADD MARGIN │ LEVERAGE │ MAX NOTIONAL VALUE │
├─────────┼─────────────┼────────────────────┼──────────┼────────────────────┤
│ BTCUSDT │ CROSSED     │ false              │ 10       │ 230000000          │
└─────────┴─────────────┴────────────────────┴──────────┴────────────────────┘
```

## Trade -  Get trades for a specific account and symbol.
Exec: `./binance-cli futures trade ls --symbol=BTCUSDT`
```shell
┌──────────────┬─────────┬──────┬───────────────┬───────────┬──────────┬────────────────┬──────────────┬─────────────────────┐
│   ORDER ID   │ SYMBOL  │ SIDE │ POSITION SIDE │   PRICE   │ QUANTITY │ QUOTE QUANTITY │ REALIZED PNL │        TIME         │
├──────────────┼─────────┼──────┼───────────────┼───────────┼──────────┼────────────────┼──────────────┼─────────────────────┤
│ 749374403589 │ BTCUSDT │ BUY  │ SHORT         │ 119711.70 │ 0.008    │ 957.69360      │ -29.22242085 │ 2025-08-11 12:51:09 │
│ 749374653742 │ BTCUSDT │ BUY  │ SHORT         │ 119722.70 │ 0.008    │ 957.78160      │ -29.31042085 │ 2025-08-11 12:51:23 │
│ 749400923511 │ BTCUSDT │ BUY  │ SHORT         │ 119907.80 │ 0.008    │ 959.26240      │ -30.79122085 │ 2025-08-11 13:36:25 │
│ 749400940699 │ BTCUSDT │ BUY  │ SHORT         │ 119887.20 │ 0.008    │ 959.09760      │ -30.62642085 │ 2025-08-11 13:36:27 │
└──────────────┴─────────┴──────┴───────────────┴───────────┴──────────┴────────────────┴──────────────┴─────────────────────┘
```
