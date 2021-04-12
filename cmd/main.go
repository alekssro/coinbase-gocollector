package main

import (
	"flag"

	"github.com/alekssro/coinbase-gocollector/pkg/collector"
	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
)

func main() {

	p := flag.String("product", "BTC-EUR", "Crypto trade value to store")
	f := flag.String("filename", "dataset.csv", "Filename to save the dataset to")
	flag.Parse()

	logger.Info("Coinbase history collector started...")
	collector.CreateDataset(*p, *f)
}
