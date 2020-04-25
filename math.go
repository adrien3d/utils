package utils

import "math"

const earthRadius = float64(6371000)

func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) +
		math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
			math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
	var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))

	distance = earthRadius * c

	return
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/*
func Distance(origin, position Point) float64 {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(origin.Lat)*math.Cos(position.Lat)*math.Pow(math.Sin(change.Lon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}*/