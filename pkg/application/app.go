package application

import (
	"log"
	"net/http"
	"time"

	"github.com/alekssro/coinbase-gocollector/pkg/util/logger"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func Start() {

	client := coinbasepro.NewClient()

	client.HTTPClient = &http.Client{
		Timeout: 15 * time.Second,
	}

	client.RetryCount = 3 // 500ms, 1500ms, 3500ms

	rates, err := client.GetHistoricRates(
		"BTC-EUR",
		coinbasepro.GetHistoricRatesParams{
			Start:       time.Now().Add(-24 * time.Hour),
			End:         time.Now(),
			Granularity: 300,
		})
	if err != nil {
		logger.Error("Error getting historic rates")
		panic(err.Error())
	}

	log.Println(rates)

}
