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
	testRecords := [][]string{{"Day", "MxT", "MnT", "AvT", "AvDP", "1HrP TPcpn", "PDir", "AvSp", "Dir", "MxS", "SkyC", "MxR", "Mn", "R AvSLP"}, {"1", "88", "59", "74", "53.8", "0", "280", "9.6", "270", "17", "1.6", "93", "23", "1004.5"}, {"2", "79", "63", "71", "46.5", "0", "330", "8.7", "340", "23", "3.3", "70", "28", "1004.5"}, {"3", "77", "55", "66", "39.6", "0", "350", "5", "350", "9", "2.8", "59", "24", "1016.8"}}
	var testRecordsEmpty [][]string
	testRecordsInt := []int{0, 59, 63, 55}
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
		{"emptySlice", args{testRecordsEmpty}, nil, true},
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
	testSliceEdge1 := []int{0, 0, 12, 3}
	testSliceEdge2 := []int{0, 1, 12, 0, 14, 7}
	type args struct {
		recordsInt []int
	}
	tests := []struct {
		name        string
		args        args
		wantDay     int
		wantMinTemp int
	}{
		{"good", args{testSlice}, 3, 4},
		{"edgeCase#1", args{testSliceEdge1}, 1, 0},
		{"edgeCase#2", args{testSliceEdge2}, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDay, gotMinTemp := GetMinTemp(tt.args.recordsInt)
			if gotDay != tt.wantDay {
				t.Errorf("GetMinTemp() gotDay = %v, want %v", gotDay, tt.wantDay)
			}
			if gotMinTemp != tt.wantMinTemp {
				t.Errorf("GetMinTemp() gotMinTemp = %v, want %v", gotMinTemp, tt.wantMinTemp)
			}
		})
	}
}

func TestSuffix(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name    string
		args    args
		wantEnd string
	}{
		{"test1st", args{1}, "st"},
		{"test2nd", args{2}, "nd"},
		{"test3rd", args{3}, "rd"},
		{"testdefault", args{120}, "th"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEnd := Suffix(tt.args.d); gotEnd != tt.wantEnd {
				t.Errorf("Suffix() = %v, want %v", gotEnd, tt.wantEnd)
			}
		})
	}
}
