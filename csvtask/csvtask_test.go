package csvtask

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOpenReadCsv(t *testing.T) {
	testCsv := [][]string{{"Day", "MxT", "MnT", "AvT", "AvDP", "1HrP TPcpn", "PDir", "AvSp", "Dir", "MxS", "SkyC", "MxR", "Mn", "R AvSLP"}, {"1", "88", "59", "74", "53.8", "0", "280", "9.6", "270", "17", "1.6", "93", "23", "1004.5"}, {"2", "79", "63", "71", "46.5", "0", "330", "8.7", "340", "23", "3.3", "70", "28", "1004.5"}, {"3", "77", "55", "66", "39.6", "0", "350", "5", "350", "9", "2.8", "59", "24", "1016.8"}}
	type args struct {
		csvfile string
	}
	tests := []struct {
		name     string
		args     args
		wantFile [][]string
		wantErr  bool
	}{
		{"good", args{"test.csv"}, testCsv, false},
		{"missingCsv", args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := OpenCsv(tt.args.csvfile)
			gotFile, err := ReadCsv(file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(gotFile, tt.wantFile); diff != "" {
				t.Errorf("test failed, diff ==> %v\n,", diff)
			}
		})
	}
}

func TestConvertCsv(t *testing.T) {
	testRecords := [][]string{{"Day", "MxT", "MnT", "AvT", "AvDP", "1HrP TPcpn", "PDir", "AvSp", "Dir", "MxS", "SkyC", "MxR", "Mn", "R AvSLP"}, {"0", "2", "6"}, {"1", "4", "5"}}
	testRecordsInt := []int{0, 6, 5}
	var testRecordsIntWrong []int
	type args struct {
		records [][]string
	}
	tests := []struct {
		name           string
		args           args
		wantRecordsInt []int
		wantErr        bool
	}{
		{"good", args{testRecords}, testRecordsInt, false},
		{"emptySlice", args{testRecords}, testRecordsIntWrong, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordsInt, err := ConvertCsv(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(gotRecordsInt, tt.wantRecordsInt); diff != "" {
				t.Errorf("test failed, diff ==> %v\n,", diff)
			}
		})
	}
}

func TestGetMinTemp(t *testing.T) {
	testSlice := []int{0, 22, 12, 4, 56, 772}
	type args struct {
		recordsInt []int
	}
	tests := []struct {
		name        string
		args        args
		wantDay     int
		wantMinTemp int
		wantErr     bool
	}{
		{"good", args{testSlice}, 3, 4, false},
		{"wrongSlice", args{testSlice}, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDay, gotMinTemp, err := GetMinTemp(tt.args.recordsInt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinTemp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDay != tt.wantDay {
				t.Errorf("GetMinTemp() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
			if gotMinTemp != tt.wantMinTemp {
				t.Errorf("GetMinTemp() gotMinTemp = %v, want %v", gotMinTemp, tt.wantMinTemp)
			}
		})
	}
}
