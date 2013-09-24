package whisper

import (
	"fmt"
)

type MetricType string

const (
	METRIC_COUNTER = "c"
)

type Metric struct {
	name      string
	value     int64
	kind      MetricType
	timestamp uint64
}

func (m Metric) String() string {
	return fmt.Sprintf("%v %v%v%v %v", m.name, m.value, METRIC_VALUE_SEP, m.kind, m.timestamp)
}
