package locations

import (
	"math"
)

type Point struct {
	Lat float64
	Lng float64
}

const EarthRadiusInMeters = 6371000 // Earth's radius in meters

// DistanceKm returns the distance in km between two points in Km.
func DistanceKm(p1, p2 Point) float64 {
	dLat := degreeToRadian(p2.Lat - p1.Lat)
	dLon := degreeToRadian(p2.Lng - p1.Lng)

	a := math.Sin(dLat/2) * math.Sin(dLat/2)
	b := math.Cos(degreeToRadian(p1.Lat)) * math.Cos(degreeToRadian(p2.Lat)) *
		math.Sin(dLon/2) * math.Sin(dLon/2)

	c := 2 * math.Asin(math.Sqrt(a+b))
	return EarthRadiusInMeters * c / 1000.0
}

func degreeToRadian(deg float64) float64 {
	return deg * (math.Pi / 180)
}
