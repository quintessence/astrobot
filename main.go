package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/qanx/astrobot/coconvert"
	"github.com/qanx/astrobot/geo"
	"github.com/qanx/astrobot/astro"
)

func main() {
	// Testing against Polaris, RA 2h31m49.09s (2.5303027778) Dec +89deg15min50.8sec (89.5141111111)

	var rightAscHours, declination float64
	flag.Float64Var(&rightAscHours, "ra", 2.5303027778, "Right ascention in decimal hours")
	flag.Float64Var(&declination, "dec", 89.5141111111, "Declination in decimal degrees")
	flag.Parse()

	//dateTime := time.Date(2015, 8, 22, 20, 32, 0, 0, time.Local)
	latitude, longitude := geo.GetCoordinates()
	//Right now this is just outputting the input
	rightAscHours, declination = astro.GetCoordinates(rightAscHours, declination)
	dateTime := time.Now()

	//Get Mac location with Objective C project here:
	//https://github.com/robmathers/WhereAmI
  //fmt.Printf("Latitude: %f\nLongitude: %f\nRA: %f\nDec: %f\n", latitude, longitude, rightAscHours, declination)
	altitudeDegrees, azimuthDegrees := coconvert.CelestialToHorizontal(rightAscHours, declination, latitude, longitude, dateTime)
	fmt.Println("The altitude in degrees is: ", altitudeDegrees)
	fmt.Println("The azimuth in degrees is: ", azimuthDegrees)
}
