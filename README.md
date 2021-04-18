# Data Collector

This is the service in charge of collecting the cryptocurrency price evolution data. It makes use of the [`go-coinbasepro`](https://pkg.go.dev/github.com/preichenberger/go-coinbasepro/v2) library to manage the Coinbase pro API and get the historic data in the form:

- `Time`: bucket start time
- `Close`: closing price (last trade) in the bucket interval
- `High`: highest price during the bucket interval
- `Low`: lowest price during the bucket interval
- `Open`: opening price (first trade) in the bucket interval
- `Volume`: volume of trading activity during the bucket interval
- `Product`: product id as of the coinbase API [/products](https://docs.pro.coinbase.com/#products)


## Example of use

Collect BTC-EUR historic data every 1 hour for the last year and save it in `out/BTC.csv` 

```
go run cmd/main.go --product="BTC-EUR" --time=365 --granularity=3600 --filename="out/BTC-EUR.csv"
```

## Usage

```
Usage of cmd/main:
  -filename string
        Filename to save the dataset to (default "dataset.csv")
  -granularity int
        Time between data points, in seconds. Options: {60, 300, 900, 3600, 21600, 86400}. (default 21600)
  -product string
        Crypto trade value to store (default "BTC-EUR")
  -time int
        Time of data collection, in days (default 30)
```

