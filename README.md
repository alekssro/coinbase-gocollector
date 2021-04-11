# Data Collector

This is the service in charge of collecting the cryptocurrency price evolution data. It makes use of the [`go-coinbasepro`](https://pkg.go.dev/github.com/preichenberger/go-coinbasepro/v2) library to manage the Coinbase pro API and get the historic data in the form:

- `time`: bucket start time
- `low`: lowest price during the bucket interval
- `high`: highest price during the bucket interval
- `open`: opening price (first trade) in the bucket interval
- `close`: closing price (last trade) in the bucket interval
- `volume`: volume of trading activity during the bucket interval
