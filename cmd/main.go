package main

import (
	"github.com/alekssro/coinbase-gocollector/pkg/application"
	"github.com/alekssro/coinbase-gocollector/pkg/util/logger"
)

func main() {
	logger.Info("Coinbase history collector started...")
	application.Start()
}
