# Binance CLI
A command-line tool for Binance API developed based on the Go language, supporting various trading functions such as spot, futures, and portfolio margin.

## Installation and Configuration

### Installation
```shell
curl -sSL https://raw.githubusercontent.com/UnipayFI/binance-cli/refs/heads/main/download.sh | bash
```

### Environment variables
Before using, you need to set the Binance API key:
```shell
export API_KEY="your_api_key"
export API_SECRET="your_api_secret"
```

## Usage
All commands are to be used in the following format:
```
./binance-cli [Module] [Subcommand] [Arguments]

Available Commands:
  futures            Futures
  help               Help about any command
  portfolio          Portfolio
  spot               Spot
  universal-transfer User universal transfer
```

### Spot Module
Exec: `./binance-cli spot [Subcommand] [Arguments]`
```shell
Available Commands:
  account     Show account info
  asset       Show account assets
  dividend    Get dividend information
  order       Support create, cancel, list orders
```
**[View detailed documentation](docs/spot.md)**


### Futures Module
Exec: `./binance-cli futures [Subcommand] [Arguments]`
```shell
Available Commands:
  account           Show account balances & account config
  commission-rate   Show commission rate
  fee               BNB payment fee
  income            Query income history
  multi-assets-mode Show and set multi-assets mode
  order             Support create, cancel, list futures orders
  position          Show positions & show position risk & set position margin & change position side
  symbol            Symbol config(leverage & margin type)
  trade             Get trades for a specific account and symbol.
```
**[View detailed documentation](docs/futures.md)**

### Portfolio Module
#### Portfolio USDⓈ-Margined Futures
Exec: `./binance-cli portfolio um [Subcommand] [Arguments]`
```shell
Available Commands:
  commission-rate Get User Commission Rate for UM
  fee             BNB payment fee
  income          Get UM Income History
  order           Support create, cancel, list um orders
  position        Show positions & show position risk & change position side
  symbol          Symbol config
```
**[View detailed documentation](docs/portfolio_um.md)**

#### Portfolio Margin Trading
Exec: `./binance-cli portfolio margin [Subcommand] [Arguments]`
```shell
Available Commands:
  interest-history Interest history
  loan             Loan
  order            Support create, cancel, list margin orders
```
**[View detailed documentation](docs/portfolio_margin.md)**

## License
UnipayFI/binance-cli is released under the [MIT License](https://opensource.org/licenses/MIT).