package hubeny

import (
	"math"
	"testing"
)

var testcases = []struct {
	lat1     float64
	lng1     float64
	lat2     float64
	lng2     float64
	distance float64
}{
	{35.689487, 139.691706, 35.5562073, 139.5723855, 18317.122},
	{36.10377477777778, 140.08785502777778, 35.65502847222223, 139.74475044444443, 58643.804},
	{34.73348, 135.500109, 35.681382, 139.766084, 402392.122},
}

func TestHubeny(t *testing.T) {
	for _, testcase := range testcases {
		result := Hubeny(
			testcase.lat1,
			testcase.lng1,
			testcase.lat2,
			testcase.lng2,
		)
		if math.Abs(1-result/testcase.distance) > 0.001 {
			t.Error("Oops", testcase.distance, result)
		}
	}
}

func BenchmarkHubeny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hubeny(
			testcases[0].lat1,
			testcases[0].lng1,
			testcases[0].lat2,
			testcases[0].lng2,
		)
	}
}
