package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {

	var infinite *bool

	infinite = flag.Bool("infi", false, "If you want to run infinite")
	initialTemp := flag.Float64("temp", 0.0, "The temperature to convert")
	originUnit := flag.String("unit", "C", "The unit of temperature")
	var repeat bool
	flag.Parse()

	fmt.Println(*infinite)
	if *infinite == false {
		Convert(originUnit, initialTemp)
	} else {
		for {
			fmt.Println("Enter the Temperature")
			_, err := fmt.Scanln(initialTemp)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Enter the Target Unit")
			_, err = fmt.Scanln(originUnit)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if len(*originUnit) != 1 {
				lenErr := errors.New("Invalid Unit")
				fmt.Println(lenErr)
				continue
			}
			Convert(originUnit, initialTemp)
			fmt.Println("Do you want to continue (true/false) ?")
			_, err = fmt.Scanln(&repeat)
			if err != nil {
				fmt.Println(err)
				break
			}
			if !repeat {
				break
			}
		}
	}
}

func Convert(originUnit *string, initialTemp *float64) {
	var temp float64
	if *originUnit == "C" {
		temp = ConvertToCelsius(*initialTemp)
	} else {
		temp = ConvertToFarhenheit(*initialTemp)
	}
	fmt.Println("New Temperature is", temp)
}

func ConvertToCelsius(value float64) float64 {
	fmt.Println("Converting", value, "to Celsuis")
	convertedValue := (value - 32) * 5 / 9
	return convertedValue
}

func ConvertToFarhenheit(value float64) float64 {
	fmt.Println("Converting", value, "to Farhenheit")
	convertedValue := (value * 9 / 5) + 32
	return convertedValue
}
