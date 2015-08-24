package main

import (
	"fmt"
	"time"

	"github.com/qanx/astrobot/coconvert"
)

func main() {
	var latitude, longitude, rightAscHours, declination float64 = 43.0, -79.0, 0.0, 0.0
	//dateTime := time.Now()
	dateTime := time.Date(2015, 8, 22, 20, 32, 0, 0, time.Local)
	//fmt.Println(latitude, longitude, rightAscHours, declination, dateTime)

	altitudeDegrees, azimuthDegrees := coconvert.CelestialToHorizon(rightAscHours, declination, latitude, longitude, dateTime)
	fmt.Println("The altitude in degrees is: ", altitudeDegrees)
	fmt.Println("The azimuth in degrees is: ", azimuthDegrees)
}
