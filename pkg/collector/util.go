package collector

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alekssro/coinbase-gocollector/pkg/shared/logger"
	"github.com/hako/durafmt"
	"github.com/preichenberger/go-coinbasepro/v2"
)

type Elapsed struct {
	duration time.Duration
}

type TimeInterval struct {
	Start time.Time
	End   time.Time
}

// Split method divides a time.Duration into s splits.
// Returns a slice with the correspondant start-end time of the splits
func (e Elapsed) Split(s int) []TimeInterval {

	remaining := e.duration
	splitDuration := time.Duration(int(e.duration.Seconds())/s) * time.Second

	var ts []TimeInterval
	for i := 0; i < s; i++ {
		start := time.Now().Add(-remaining)
		next := start.Add(splitDuration)
		fmt.Println(remaining)
		fmt.Println(start, next)
		ts = append(ts, TimeInterval{
			Start: start,
			End:   next,
		})
		remaining = remaining - time.Duration(splitDuration)
	}

	return ts
}

func (e Elapsed) String() string {
	return durafmt.Parse(e.duration).String()
}

func getClient() coinbasepro.Client {
	client := coinbasepro.NewClient()
	client.RetryCount = 3 // 500ms, 1500ms, 3500ms
	client.HTTPClient = &http.Client{
		Timeout: 15 * time.Second,
	}
	return *client
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

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
