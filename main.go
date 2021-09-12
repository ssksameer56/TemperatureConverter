package main

import (
	"flag"
	"fmt"
)

func main() {
	initialTemp := flag.Float64("temp", 0.0, "The temperature to convert")
	originUnit := flag.String("unit", "C", "The unit of temperature")

	flag.Parse()

	var temp float64
	if *originUnit == "C" {
		temp = ConvertToCelsius(*initialTemp)
	} else {
		temp = ConvertToFarhenheit(*initialTemp)
	}
	fmt.Println("New Temperature is ", temp)
}

func ConvertToCelsius(value float64) float64 {
	fmt.Println("Converting to Celsuis", value)
	convertedValue := (value - 32) * 5 / 9
	return convertedValue
}

func ConvertToFarhenheit(value float64) float64 {
	fmt.Println("Converting to Farhenheit", value)
	convertedValue := (value * 9 / 5) + 32
	return convertedValue
}
