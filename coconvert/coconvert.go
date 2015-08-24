package coconvert

import (
	"math"
	"time"
)

// 1HR RA = 15 degrees
func ra2deg(raHr float64) float64 {
	return raHr * 15
}

// Returns mean sidereal time in degrees. Based on p89 Astronomical Algorithms 2nd Ed by John Meeus
func meanSiderealTime(long float64, dateTime time.Time) float64 {
	year := float64(dateTime.Year())
	month := float64(dateTime.Month())
	day := float64(dateTime.Day())
	hour := float64(dateTime.Hour())
	minute := float64(dateTime.Minute())
	second := float64(dateTime.Second())

	if month == 1 || month == 2 {
		month = month + 12
		year = year - 1
	}

	//Constants
	a := math.Floor(year / 100.0)
	b := 2 - a + math.Floor(a/4.0)
	c := math.Floor(365.25 * year)
	d := math.Floor(30.6001 * (month + 1))

	//Days since J2000.0
	jd := b + c + d - 730550.5 + day + (hour+minute/60.0+second/3600.0)/24.0

	//Centuries since J2000.0
	jt := jd / 36525.0

	meansidereal := 280.46061837 + 360.98564736629*jd + 0.000387933*jt*jt - jt*jt*jt/38710000 + long

	//Since in degrees, result range 0..360
	if meansidereal > 0 {
		for meansidereal > 360 {
			meansidereal = meansidereal - 360
		}
	} else {
		for meansidereal < 0 {
			meansidereal = meansidereal + 360
		}
	}

	return meansidereal
}

func CelestialToHorizontal(raHr float64, dec float64, lat float64, long float64, dateTime time.Time) (float64, float64) {
	radeg := ra2deg(raHr)
	hourAngle := meanSiderealTime(long, dateTime) - radeg
	if hourAngle < 0 {
		hourAngle = hourAngle + 360
	}

	hourAngle = hourAngle * math.Pi / 180
	dec = dec * math.Pi / 180
	lat = lat * math.Pi / 180

	//azimuth error: divide by zero error at poles or if alt = 90 deg
	altitudeRad := math.Asin(math.Sin(dec)*math.Sin(lat) + math.Cos(dec)*math.Cos(lat)*math.Cos(hourAngle))
	azimuthRad := math.Acos((math.Sin(dec) - math.Sin(altitudeRad)*math.Sin(lat)) / (math.Cos(altitudeRad) * math.Cos(lat)))

	return altitudeRad * 180 / math.Pi, azimuthRad * 180 / math.Pi
}
