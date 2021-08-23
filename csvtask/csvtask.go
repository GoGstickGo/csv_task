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
	for _, v := range records {
		data := csvData{
			MinTempRow: v[2],
		}
		//convert string elements to integer
		tempInt, _ := strconv.Atoi(data.MinTempRow)

		recordsInt = append(recordsInt, tempInt)
		if len(recordsInt) == 0 {
			return nil, fmt.Errorf("can't proceed with empty slice")
		}
	}
	return recordsInt, nil
}

func GetMinTemp(recordsInt []int) (day, minTemp int, err error) {
	//remove first element of slice
	recordsInt = append(recordsInt[:0], recordsInt[0+1:]...)
	if recordsInt[0] == 0 {
		return 0, 0, fmt.Errorf("couldn't get lowest temperature")
	}
	minTemp = recordsInt[0]
	for k, v := range recordsInt {
		if v < minTemp {
			minTemp = v
			day = k + 1
		}
	}
	return day, minTemp, nil
}
