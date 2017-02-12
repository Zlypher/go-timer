package interval

import "testing"

func TestParseInterval(t *testing.T) {
	cases := []struct {
		in   []string
		want []int64
	}{
		{[]string{}, []int64{}},
		{[]string{"100ms"}, []int64{100}},
		{[]string{"1s300ms"}, []int64{1300}},
		{[]string{"5m12s"}, []int64{312000}},
	}

	for _, c := range cases {
		result := parseArguments(c.in)
		intervals := result.intervals

		if len(intervals) != len(c.want) {
			t.Errorf("len(parseArguments(%v).intervals) == %d, want %d", c.in, len(intervals), len(c.want))
		}

		for i, w := range c.want {
			got := intervals[i].Nanoseconds() / 1e6
			if got != w {
				t.Errorf("Interval(%d) == %d, want %d", i, got, w)
			}
		}
	}
}
