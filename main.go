package main

import (
	"fmt"
	"log"
	"weather-temps/csvtask"
)

func main() {
	file, err := csvtask.OpenCsv("weather.csv")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	records, err := csvtask.ReadCsv(file)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	recordsInt, err := csvtask.ConvertCsv(records)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	day, minTemp, err := csvtask.GetMinTemp(recordsInt)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Lowest temperature was on %dth day, temperature was %d Fahrenheit.\n", day, minTemp)

}
