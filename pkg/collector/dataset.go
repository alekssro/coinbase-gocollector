package collector

import "time"

// Granularity options
// {60, 300, 900, 3600, 21600, 86400}
// {1min, 5mins, 15 mins, 1 hour, 6 hours, 1 day}

type DatasetArgs struct {
	Product     *string
	Filename    *string
	Time        *int // in days
	Granularity *int // in secs
}

type DatasetConfig struct {
	Product     string
	Filename    string
	StartDate   time.Time
	EndDate     time.Time
	Granularity int
}

func (a DatasetArgs) ToConfig() DatasetConfig {
	return DatasetConfig{
		Product:     *a.Product,
		Filename:    *a.Filename,
		StartDate:   time.Now().Add(-time.Hour * 24 * time.Duration(*a.Time)),
		EndDate:     time.Now(),
		Granularity: *a.Granularity,
	}
}
