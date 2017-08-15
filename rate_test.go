package buildinfo

import "testing"

func TestRate(t *testing.T) {
	cases := []struct {
		in   BuildInfo
		want float32
	}{
		{BuildInfo{0.3, 0, "link", 0.7}, 0.75},
		{BuildInfo{0.3, 0, "", 0.7}, 0.5},
		{BuildInfo{0.3, 0.7, "link", 0.7}, 1},
	}
	for _, c := range cases {
		got := checkRating(c.in)
		if got != c.want {
			t.Errorf("checkRating(%v) == %f, want %f", c.in, got, c.want)
		}
	}
}
