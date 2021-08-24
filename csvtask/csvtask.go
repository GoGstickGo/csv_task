package csvtask

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type csvData struct {
	MinTempRow string
}

func OpenCsv(csvfile string) (file *os.File, err error) {
	file, err = os.Open(csvfile)
	if err != nil {
		return nil, fmt.Errorf("couldn't open the csv file: %v", err)
	}
	return file, nil
}

func ReadCsv(file *os.File) (records [][]string, err error) {
	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("csv reader failed: %v", err)
	}
	return records, nil
}

func ConvertCsv(records [][]string) (recordsInt []int, err error) {
	if len(records) == 0 {
		return nil, fmt.Errorf("there is no data to process, please make csv file has data")
	}
	for _, v := range records {
		data := csvData{
			MinTempRow: v[2],
		}
		//convert string elements to integer
		tempInt, _ := strconv.Atoi(data.MinTempRow)

		recordsInt = append(recordsInt, tempInt)
	}
	return recordsInt, nil
}

func GetMinTemp(recordsInt []int) (day, minTemp int) {
	//remove first element of slice
	recordsInt = append(recordsInt[:0], recordsInt[0+1:]...)

	for k, v := range recordsInt {
		if k == 0 || v < minTemp {
			minTemp = v
			day = k + 1
		}
	}
	return day, minTemp
}

func Suffix(d int) (end string) {
	switch d {
	case 1:
		end = "st"
	case 2:
		end = "nd"
	case 3:
		end = "rd"
	default:
		end = "th"
	}
	return end
}
