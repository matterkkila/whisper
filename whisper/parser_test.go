package whisper

import (
	"testing"
)

func BenchmarkParseMetric(b *testing.B) {
	s := []byte("my.metric.name 1|c 1379700446")
	p := TextParser{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestParseMetric(t *testing.T) {
	s := "my.metric.name 1|c 1379700446"
	p := TextParser{}
	m, err := p.Parse([]byte(s))
	if err != nil {
		t.Error(err)
	}
	if m.name != "my.metric.name" {
		t.Error("Invalid name")
	}
	if m.value != 1 {
		t.Error("Invalid value")
	}
	if m.kind != METRIC_COUNTER {
		t.Error("Invalid kind")
	}
	if m.timestamp != 1379700446 {
		t.Error("Invalid ts")
	}
	if m.String() != "my.metric.name 1|c 1379700446" {
		t.Error("Invalid string conversion.", m.String())
	}
}
