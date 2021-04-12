package collector

import (
	"net/http"
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
)

func getClient() coinbasepro.Client {
	client := coinbasepro.NewClient()
	client.RetryCount = 3 // 500ms, 1500ms, 3500ms
	client.HTTPClient = &http.Client{
		Timeout: 15 * time.Second,
	}
	return *client
}
