package wug

import (
	"os"
	"testing"
)

var testApiKey string

func init() {
	testApiKey = os.Getenv("WUGKEY")
}

func TestRawConditions(t *testing.T) {
	wug := NewWug()
	q := NewQueryByAutoIP(testApiKey)
	data, err := wug.GetRawConditions(q)
	if err != nil {
		t.Fatalf("error getting conditions: %s\n", err)
	}

	if len(data) == 0 {
		t.Fatalf("error data was empty")
	}

	t.Logf("%s\n", string(data))
}

func TestRawForecast(t *testing.T) {
	wug := NewWug()
	q := NewQueryByAutoIP(testApiKey)
	data, err := wug.GetRawForecast(q)
	if err != nil {
		t.Fatalf("error getting conditions: %s\n", err)
	}

	if len(data) == 0 {
		t.Fatalf("error data was empty")
	}

	t.Logf("%s\n", string(data))
}

func TestForecast(t *testing.T) {
	wug := NewWug()
	q := NewQueryByAutoIP(testApiKey)
	data, err := wug.GetForecast(q)
	if err != nil {
		t.Fatalf("error getting forecast: %s\n", err)
	}

	t.Logf("%#v\n", data)
}

func TestHourly(t *testing.T) {
	wug := NewWug()
	q := NewQueryByAutoIP(testApiKey)
	data, err := wug.GetHourly(q)
	if err != nil {
		t.Fatalf("error getting hourly: %s\n", err)
	}

	t.Logf("%#v\n", data)
}

func TestHourlyTenDay(t *testing.T) {
	wug := NewWug()
	q := NewQueryByAutoIP(testApiKey)
	data, err := wug.GetHourlyTenDay(q)
	if err != nil {
		t.Fatalf("error getting hourly: %s\n", err)
	}

	t.Logf("%#v\n", data.Hourly)
}

func TestHourlyTenDayLatLong(t *testing.T) {
	wug := NewWug()
	q := NewQueryByLatLong(testApiKey, "35.350178", "139.623993")
	_, err := wug.GetHourlyTenDay(q)
	if err != nil {
		t.Fatalf("error getting hourly: %s\n", err)
	}
}
