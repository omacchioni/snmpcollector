package output

import (
	"sync"
	"time"
)

// InfluxStats  get stats
type InfluxStats struct {
	// Points Sent
	PSent int64
	// PSentMax the max
	PSentMax int64
	// WriteSent BatchPoints sent
	WriteSent int64
	// WriteErrors BatchPoints with  errors
	WriteErrors int64
	// WriteTime
	WriteTime time.Duration
	// WriteTimeMax
	WriteTimeMax time.Duration
	mutex        sync.Mutex
}

// GetResetStats get stats for this InfluxStats Output
func (is *InfluxStats) GetResetStats() *InfluxStats {
	is.mutex.Lock()
	defer is.mutex.Unlock()
	retstat := &InfluxStats{
		PSent:        is.PSent,
		PSentMax:     is.PSentMax,
		WriteSent:    is.WriteSent,
		WriteErrors:  is.WriteErrors,
		WriteTime:    is.WriteTime,
		WriteTimeMax: is.WriteTimeMax,
	}
	is.PSent = 0
	is.PSentMax = 0
	is.WriteSent = 0
	is.WriteErrors = 0
	is.WriteTime = 0
	is.WriteTimeMax = 0
	return retstat
}

// WriteOkUpdate update stats on write ok
func (is *InfluxStats) WriteOkUpdate(ps int64, wt time.Duration) {
	is.mutex.Lock()
	defer is.mutex.Unlock()
	if is.PSentMax < ps {
		is.PSentMax = ps
	}
	if is.WriteTimeMax < wt {
		is.WriteTimeMax = wt
	}
	is.WriteSent++
	is.PSent += ps
	is.WriteTime += wt
}

// WriteErrUpdate update stats on write error
func (is *InfluxStats) WriteErrUpdate(wt time.Duration) {
	is.mutex.Lock()
	defer is.mutex.Unlock()

	if is.WriteTimeMax < wt {
		is.WriteTimeMax = wt
	}
	is.WriteErrors++
	is.WriteTime += wt
}
