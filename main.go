package main

import (
	"fmt"
	"time"

	"github.com/qanx/astrobot/coconvert"
)

func main() {
	// Testing against Polaris, RA 2h31m49.09s (2.5303027778) Dec +89deg15min50.8sec (89.5141111111)
	var latitude, longitude, rightAscHours, declination float64 = 43.0, -79.0, 2.5303027778, 89.5141111111
	dateTime := time.Date(2015, 8, 22, 20, 32, 0, 0, time.Local)
	//dateTime := time.Now()
	//fmt.Println(latitude, longitude, rightAscHours, declination, dateTime)

	altitudeDegrees, azimuthDegrees := coconvert.CelestialToHorizontal(rightAscHours, declination, latitude, longitude, dateTime)
	fmt.Println("The altitude in degrees is: ", altitudeDegrees)
	fmt.Println("The azimuth in degrees is: ", azimuthDegrees)
}
