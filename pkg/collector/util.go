package collector

import (
	"net/http"
	"strconv"
	"time"

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

// ByDate implements sort.Interface for []coinbasepro.HistoricRate
// based on the Time field.
type ByDate []coinbasepro.HistoricRate

func (r ByDate) Len() int           { return len(r) }
func (r ByDate) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByDate) Less(i, j int) bool { return r[i].Time.Before(r[j].Time) }

// Split method divides a time.Duration into s splits.
// Returns a slice with the correspondant start-end time of the splits
func (e Elapsed) Split(s int) []TimeInterval {

	remaining := e.duration
	splitDuration := time.Duration(int(e.duration.Seconds())/s) * time.Second

	var ts []TimeInterval
	for i := 0; i < s; i++ {
		start := time.Now().Add(-remaining)
		next := start.Add(splitDuration)
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

func IntToStr(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
