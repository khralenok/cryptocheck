# Cryptocheck(In production)

**Cryptocheck** is a fast and simple command-line interface (CLI) tool for checking real-time cryptocurrency prices and discovering the top assets by market capitalization.

## Features

- 🔎 Check the current price of any cryptocurrency in various fiat currencies (e.g. USD, EUR, JPY).
- 📈 Display a list of cryptocurrencies ranked by highest market capitalization.
- 🌐 Uses live data from a CoinCap API .

## Installation

### Go install (recommended)

If you have Go installed:

```
go install github.com/khralenok/cryptocheck@latest
git clone https://github.com/khralenok/cryptocheck.git
cd cryptocheck
go build -o cryptocheck
./cryptocheck
```

### From source

```
git clone https://github.com/khralenok/cryptocheck.git
cd cryptocheck
go build -o cryptocheck
./cryptocheck
```

## Usage

### Check price of a specific cryptocurrency

```
cryptocheck price bitcoin --fiat usd
```

### Output:

```
Bitcoin (BTC): $63,500.42 USD
```

You can change bitcoin to eth, doge, etc., and set --fiat to any supported currency (e.g. eur, jpy, aud).

### Show top cryptocurrencies by market cap

```
cryptocheck top --limit 10
```

### Output:

```
1. Bitcoin (BTC): $1.2T
2. Ethereum (ETH): $500B
3. Tether (USDT): $100B
...
```

## Dependencies

1. Go >= 1.20
2. CoinCap
3. Frankfurter

## License

MIT License. See LICENSE for details.
