package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/qanx/astrobot/coconvert"
)

func main() {
	// Testing against Polaris, RA 2h31m49.09s (2.5303027778) Dec +89deg15min50.8sec (89.5141111111)
	//var latitude, longitude, rightAscHours, declination float64 = 43.0, -79.0, 2.5303027778, 89.5141111111
	var latitude, longitude, rightAscHours, declination float64
	flag.Float64Var(&latitude, "lat", 43.0, "Latitude in degrees")
	flag.Float64Var(&longitude, "long", -79.0, "Longitude in degrees")
	flag.Float64Var(&rightAscHours, "ra", 2.5303027778, "Right ascention in decimal hours")
	flag.Float64Var(&declination, "dec", 89.5141111111, "Declination in decimal degrees")
	flag.Parse()
	//fmt.Printf("Latitude: %f\nLongitude: %f\nRA: %f\nDec: %f\n", latitude, longitude, rightAscHours, declination)
	//dateTime := time.Date(2015, 8, 22, 20, 32, 0, 0, time.Local)
	dateTime := time.Now()
	//fmt.Println(latitude, longitude, rightAscHours, declination, dateTime)
	//Get Mac location with Objective C project here:
	//https://github.com/robmathers/WhereAmI

	altitudeDegrees, azimuthDegrees := coconvert.CelestialToHorizontal(rightAscHours, declination, latitude, longitude, dateTime)
	fmt.Println("The altitude in degrees is: ", altitudeDegrees)
	fmt.Println("The azimuth in degrees is: ", azimuthDegrees)
}
