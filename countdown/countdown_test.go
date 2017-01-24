package countdown

import (
	"testing"
	"time"
)

func TestParseDurationSanity(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"100ms", 100},
		{"500ms", 500},
		{"1s", 1000},
		{"1m", 60000},
		{"10m", 600000},
		{"1h", 3600000},
	}

	for _, c := range cases {
		result, err := time.ParseDuration(c.in)
		if err != nil {
			t.Error(err)
		}

		got := result.Nanoseconds() / 1e6
		if got != c.want {
			t.Errorf("ParseDuration(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
