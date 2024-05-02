package main

import (
	"reflect"
	"testing"
)

func Test_mapToJson(t *testing.T) {
	gotJson, err := mapToJson()
	if err != nil {
		t.Fatal(err)
	}

	if gotJson == nil {
		t.Fatal("returned json is a nil value")
	}

	if len(gotJson) <= 0 {
		t.Fatal("reurned json value is 0 len byte array")
	}
}

func BenchmarkStructToJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structToJson()
	}
}

func BenchmarkMapToJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapToJson()
	}
}

func BenchmarkStructToJson10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structToJson10()
	}
}

func BenchmarkMapToJson10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mapToJson10()
	}
}

func Test_structToJson(t *testing.T) {
	gotJson, err := structToJson()
	if err != nil {
		t.Fatal(err)
	}

	if gotJson == nil {
		t.Fatal("returned json is a nil value")
	}

	if len(gotJson) <= 0 {
		t.Fatal("reurned json value is 0 len byte array")
	}
}

func Test_tableMapToJson(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mapToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("mapToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
