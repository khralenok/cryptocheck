# Cryptocheck

**Cryptocheck** is a fast and simple command-line interface (CLI) tool for checking real-time cryptocurrency prices and discovering the top assets by market capitalization.

## Features

- 🔎 Check the current price of any cryptocurrency in various fiat currencies (e.g. USD, EUR, JPY).
- 📈 Display a list of cryptocurrencies ranked by highest market capitalization(From 1 to 100).
- 🌐 Uses live data from a CoinDesk API.

## Installation

```
git clone https://github.com/khralenok/cryptocheck.git
cd cryptocheck
go build -o cryptocheck
./cryptocheck help
```

## Usage

### Check price of a specific cryptocurrency

```
cryptocheck check -symbols="BTC, ETH, XRP" -fiat="EUR"
```

### Output:

```
Ethereum(ETH): 2362.6291 EUR
XRP(XRP): 2.2841 EUR
Bitcoin(BTC): 92590.5492 EUR
```

At least one value for -symbol flag is required
By default -fiat value is "USD" so you can ommit it.
You can use multiple symbols but only one fiat currency.
If some of asked symbols doesn't exitst you still will get price for the rest.
If such fiat currency doesn't exist you will get corresponding error

### Show top cryptocurrencies by market cap

```
cryptocheck top -amount="5" -fiat="VND"
```

### Output:

```
1. Bitcoin (BTC): 53380537784237232.00 VND
2. Ethereum (ETH): 8277572964681899.00 VND
3. Tether (USDT): 3901733355160805.50 VND
4. Xrp (XRP): 3881918114425700.00 VND
5. Solana (SOL): 2426890499138538.50 VND
```

By default -amount value is "10" so you can ommit it.
By default -fiat value is "USD" so you can ommit it.
You can use only one fiat currency at once.
If such fiat currency doesn't exist you will get corresponding error

## Dependencies

1. Go >= 1.20
2. CoinDesk API v2

## License

MIT License. See LICENSE for details.
