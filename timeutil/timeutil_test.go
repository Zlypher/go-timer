package timeutil

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"100ms", "00:00:00.100"},
		{"500ms", "00:00:00.500"},
		{"1s", "00:00:01.000"},
		{"1m", "00:01:00.000"},
		{"10m", "00:10:00.000"},
		{"1h", "01:00:00.000"},
		{"12h34m56s789ms", "12:34:56.789"},
	}

	for _, c := range cases {
		duration, err := time.ParseDuration(c.in)
		if err != nil {
			t.Error(err)
		}

		got := Print(duration)
		if got != c.want {
			t.Errorf("ParseDuration(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
