package main

import (
	"flag"

	"github.com/alekssro/coinbase-gocollector/pkg/collector"
	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
)

func main() {

	p := flag.String("product", "BTC-EUR", "Crypto trade value to store")
	f := flag.String("filename", "dataset.csv", "Filename to save the dataset to")
	t := flag.Int("time", 30, "Time of data collection, in days")
	g := flag.Int("granularity", 21600, "Time between data points, in seconds")
	flag.Parse()

	logger.Info("Coinbase history collector started...")

	datasetArgs := collector.DatasetArgs{
		Product:     p,
		Filename:    f,
		Time:        t,
		Granularity: g,
	}

	collector.CreateDataset(datasetArgs.ToConfig())
}
