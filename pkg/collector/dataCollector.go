package collector

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func CreateDataset(p string, f string) {

	client := getClient()

	// query values
	product := p
	startDate := time.Now().Add(-30 * 24 * time.Hour) // 1 month
	endDate := time.Now()
	granularity := 21600 // every 6 hours
	// {60, 300, 900, 3600, 21600, 86400}
	// {1min, 5mins, 15 mins, 1 hour, 6 hours, 1 day}

	rates, err := client.GetHistoricRates(
		product,
		coinbasepro.GetHistoricRatesParams{
			Start:       startDate,
			End:         endDate,
			Granularity: granularity,
		})
	if err != nil {
		logger.Error("Error getting historic rates")
		panic(err.Error())
	}

	SaveRates(rates, f)

	logger.Info("Rates saved to Dataset")
}

func SaveRates(rates []coinbasepro.HistoricRate, f string) error {

	ToCSV(rates, f)

	return nil
}

func ToCSV(rates []coinbasepro.HistoricRate, f string) error {
	file, err := os.Create(f)
	if err != nil {
		logger.Error("Cannot create file" + err.Error())
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write csv header
	header := []string{"Time", "Close", "High", "Low", "Open", "Volume"}
	writer.Write(header)

	var row []string
	for _, r := range rates {
		row = []string{
			r.Time.String(),
			fmt.Sprintf("%f", r.Close),
			fmt.Sprintf("%f", r.High),
			fmt.Sprintf("%f", r.Low),
			fmt.Sprintf("%f", r.Open),
			fmt.Sprintf("%f", r.Volume),
		}

		err := writer.Write(row)
		if err != nil {
			logger.Error("Cannot write to file" + err.Error())
			return err
		}
	}
	return nil
}
