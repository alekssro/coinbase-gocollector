package collector

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/schollz/progressbar/v3"
)

func CreateDataset(c DatasetConfig) error {

	client := getClient()

	requests := c.ToValidRequests()

	var HistoricRates []coinbasepro.HistoricRate
	var rates []coinbasepro.HistoricRate
	var err error
	// create and start new bar
	bar := progressbar.Default(int64(len(requests)), "Requesting")
	for _, req := range requests {
		bar.Add(1)
		rates, err = client.GetHistoricRates(
			c.Product,
			req,
		)
		if err != nil {
			logger.Error("Error getting historic rates: " + err.Error())
			return err
		}

		sort.Sort(ByDate(rates))

		HistoricRates = append(HistoricRates, rates...)
	}

	logger.Info(IntToStr(len(HistoricRates)) + " rates obtained")

	SaveRates(HistoricRates, c.Filename, c.Product)

	logger.Info("Rates saved to " + c.Filename)

	return nil
}

// ToValidRequests method takes the desired DatasetConfig and translate it
// to an array of valid requests for Coinbase Pro API:
// 		(Count of aggregations requested exceeds 300)
func (c DatasetConfig) ToValidRequests() []coinbasepro.GetHistoricRatesParams {

	elapsed := Elapsed{c.EndDate.Sub(c.StartDate)}
	logger.Info("Elapsed time requested: " + elapsed.String())

	granularity := Elapsed{time.Duration(c.Granularity * int(time.Second))}
	logger.Info("Granularity requested: " + granularity.String())

	points := int(elapsed.duration.Seconds()) / c.Granularity
	logger.Info("Resulting data points: " + IntToStr(points))

	splits := points/300 + 1
	logger.Info("Getting data in " + IntToStr(splits) + " requests")

	timeIntervals := elapsed.Split(splits)

	var reqParams []coinbasepro.GetHistoricRatesParams
	for _, ti := range timeIntervals {
		reqParams = append(reqParams, coinbasepro.GetHistoricRatesParams{
			Start:       ti.Start,
			End:         ti.End,
			Granularity: c.Granularity,
		})
	}

	return reqParams

}

func SaveRates(rates []coinbasepro.HistoricRate, f string, p string) error {

	ToCSV(rates, f, p)

	return nil
}

func ToCSV(rates []coinbasepro.HistoricRate, f string, p string) error {
	file, err := os.Create(f)
	if err != nil {
		logger.Error("Cannot create file" + err.Error())
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write csv header
	header := []string{"Time", "Close", "High", "Low", "Open", "Volume", "Product"}
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
			p,
		}

		err := writer.Write(row)
		if err != nil {
			logger.Error("Cannot write to file" + err.Error())
			return err
		}
	}
	return nil
}
