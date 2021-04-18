package collector

import (
	"time"

	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func CreateDataset(c DatasetConfig) error {

	client := getClient()

	requests := c.ToValidRequests()

	var HistoricRates []coinbasepro.HistoricRate
	var rates []coinbasepro.HistoricRate
	var err error
	for _, req := range requests {
		rates, err = client.GetHistoricRates(
			c.Product,
			req,
		)
		if err != nil {
			logger.Error("Error getting historic rates: " + err.Error())
			return err
		}

		HistoricRates = append(HistoricRates, rates...)
	}

	logger.Info(IntToStr(len(HistoricRates)) + " rates obtained")

	SaveRates(HistoricRates, c.Filename)

	logger.Info("Rates saved to Dataset")

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

func SaveRates(rates []coinbasepro.HistoricRate, f string) error {

	ToCSV(rates, f)

	return nil
}
