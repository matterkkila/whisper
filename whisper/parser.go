package whisper

import (
	"errors"
	"strconv"
	"strings"
)

const (
	METRIC_SEP       = " "
	METRIC_VALUE_SEP = "|"
)

type Parser interface {
	Parse([]byte) (*Metric, error)
}

type TextParser struct {}

func (p TextParser) Parse(buf []byte) (*Metric, error) {
	txt := string(buf)

	splitText := strings.SplitN(txt, METRIC_SEP, 3)
	if len(splitText) != 3 {
		return nil, errors.New("Not enough pieces in metric")
	}
	value := strings.SplitN(splitText[1], METRIC_VALUE_SEP, 2)
	if len(value) != 2 {
		return nil, errors.New("Not enough pieces in value")
	}
	intValue, err := strconv.ParseInt(value[0], 10, 64)
	if err != nil {
		return nil, err
	}

	intTimestamp, err := strconv.ParseUint(splitText[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Metric{splitText[0], intValue, MetricType(value[1]), intTimestamp}, nil
}
