# Wallet Module

## Quick Navigation
- [Dust](#dust)
  - [Get assets that can be converted into BNB](#dust---get-assets-that-can-be-converted-into-bnb)
  - [Convert dust assets to BNB](#dust---convert)
  - [Dust conversion history](#dust---history)
- [Fee](#fee)
  - [Get BNB Burn Status](#fee---get-bnb-burn-status)
  - [Set BNB Burn Status](#fee---set-bnb-burn-status)
- [Universal Transfer](#universal-transfer)
  - [Transfer between accounts](#universal-transfer---transfer)
  - [Transfer history](#universal-transfer---history)

## Dust
### Dust - Get assets that can be converted into BNB
Exec: `./binance-cli wallet dust show`
```shell
┌────────────┬────────────┬──────────────┬───────────────────────┬────────────────────┬─────────────────────┐
│ FROM ASSET │   AMOUNT   │ TRANSFER ID  │ SERVICE CHARGE AMOUNT │ TRANSFERRED AMOUNT │    OPERATE TIME     │
├────────────┼────────────┼──────────────┼───────────────────────┼────────────────────┼─────────────────────┤
│ BFUSD      │ 0.4254     │ 327477161221 │ 0.00000923            │ 0.00046158         │ 2025-11-19 06:21:09 │
│ ETH        │ 0.00003919 │ 327477161221 │ 0.00000257            │ 0.00012861         │ 2025-11-19 06:21:09 │
└────────────┴────────────┴──────────────┴───────────────────────┴────────────────────┴─────────────────────┘

Total Service Charge: 0.0000118, Total Transfered: 0.00059019
```

### Dust - Convert
Exec: `./binance-cli wallet dust convert --asset=BFUSD,ETH`
```shell
┌────────────┬────────────┬──────────────┬───────────────────────┬────────────────────┬─────────────────────┐
│ FROM ASSET │   AMOUNT   │ TRANSFER ID  │ SERVICE CHARGE AMOUNT │ TRANSFERRED AMOUNT │    OPERATE TIME     │
├────────────┼────────────┼──────────────┼───────────────────────┼────────────────────┼─────────────────────┤
│ BFUSD      │ 0.4254     │ 327477161221 │ 0.00000923            │ 0.00046158         │ 2025-11-19 06:21:09 │
│ ETH        │ 0.00003919 │ 327477161221 │ 0.00000257            │ 0.00012861         │ 2025-11-19 06:21:09 │
└────────────┴────────────┴──────────────┴───────────────────────┴────────────────────┴─────────────────────┘

Total Service Charge: 0.0000118, Total Transfered: 0.00059019
```

### Dust - History
Exec: `./binance-cli wallet dust history`
```shell
┌────────────┬────────────┬──────────────┬───────────────────────┬────────────────────┬─────────────────────┐
│ FROM ASSET │   AMOUNT   │ TRANSFER ID  │ SERVICE CHARGE AMOUNT │ TRANSFERRED AMOUNT │    OPERATE TIME     │
├────────────┼────────────┼──────────────┼───────────────────────┼────────────────────┼─────────────────────┤
│ BFUSD      │ 0.4254     │ 327477161221 │ 0.00000923            │ 0.00046158         │ 2025-11-19 06:21:10 │
│ ETH        │ 0.00003919 │ 327477161221 │ 0.00000257            │ 0.00012861         │ 2025-11-19 06:21:10 │
└────────────┴────────────┴──────────────┴───────────────────────┴────────────────────┴─────────────────────┘
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

## Universal-transfer
### Universal-transfer - transfer
Exec: `./binance-cli wallet universal-transfer --type={type} --asset={asset} --amount={x}`

### Universal-transfer - History
Exec: `./binance-cli universal-transfer ls --type={type}`
```shell
┌───────┬────────┬───────────────────────┬───────────┬──────────────┬─────────────────────┐
│ ASSET │ AMOUNT │         TYPE          │  STATUS   │   TRAN ID    │      TIMESTAMP      │
├───────┼────────┼───────────────────────┼───────────┼──────────────┼─────────────────────┤
│ USDC  │ 400000 │ MAIN_PORTFOLIO_MARGIN │ CONFIRMED │ 324539848979 │ 2025-11-13 03:37:24 │
└───────┴────────┴───────────────────────┴───────────┴──────────────┴─────────────────────┘
```