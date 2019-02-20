package hubeny

import "math"

// Hubeny returns distance between latlng1 and latlng2
func Hubeny(lat1, lng1, lat2, lng2 float64) float64 {

	radLat1 := lat1 * math.Pi / 180
	radLng1 := lng1 * math.Pi / 180
	radLat2 := lat2 * math.Pi / 180
	radLng2 := lng2 * math.Pi / 180

	radLatDiff := radLat1 - radLat2

	radLngDiff := radLng1 - radLng2

	radLatAve := (radLat1 + radLat2) / 2.0

	sinLat := math.Sin(radLatAve)

	sinLat = math.Sin(radLatAve)
	w2 := 1.0 - 0.00669438002301188*(sinLat*sinLat)
	m := 6335439.32708317 / (math.Sqrt(w2) * w2)
	n := 6378137.0 / math.Sqrt(w2)

	t1 := m * radLatDiff
	t2 := n * math.Cos(radLatAve) * radLngDiff
	dist := math.Sqrt((t1 * t1) + (t2 * t2))

	return dist
}
