package wug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// The weather underground API format string
const requestURL = "http://api.wunderground.com/api/%s/%s/q/%s"

// RequestType that is supported for weather underground requests
type RequestType int

// RequestType constants
const (
	Cond       RequestType = iota // Conditions request
	Fore                          // Forecast request
	ForeTenDay                    // Ten Day Forecast request
	Hour                          // Hourly request
	HourTenDay                    // Ten Day Hourly request
)

var requestMap = map[RequestType]string{
	Cond:       "conditions",
	Fore:       "forecast",
	ForeTenDay: "forecast10day",
	Hour:       "hourly",
	HourTenDay: "hourly10day",
}

// Wug API client that uses Query's to request data from weather underground
type Wug struct {
	Client *http.Client
}

// NewWug returns a Wug client with a configured http.Client and transport.
func NewWug() *Wug {
	var tr = &http.Transport{
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
	}
	var client = &http.Client{Transport: tr}
	return &Wug{Client: client}
}

// GetRawConditions returns the raw bytes of a conditions request
func (w *Wug) GetRawConditions(query *Query) ([]byte, error) {
	return w.Get(Cond, query)
}

// GetConditions returns the Conditions of the request
func (w *Wug) GetConditions(query *Query) (*Conditions, error) {
	data, err := w.GetRawConditions(query)
	if err != nil {
		return nil, err
	}

	conditions := &Conditions{}
	err = json.Unmarshal(data, conditions)
	if err != nil {
		return nil, err
	}
	return conditions, nil
}

// GetRawHourly returns the raw bytes of an hourly request
func (w *Wug) GetRawHourly(query *Query) ([]byte, error) {
	return w.Get(Hour, query)
}

// GetHourly returns the Hourly data.
func (w *Wug) GetHourly(query *Query) (*Hourly, error) {
	data, err := w.GetRawHourly(query)
	if err != nil {
		return nil, err
	}

	hourly := &Hourly{}
	err = json.Unmarshal(data, hourly)
	if err != nil {
		return nil, err
	}
	return hourly, nil
}

// GetRawHourlyTenDay returns the raw bytes of an hourly ten day request
func (w *Wug) GetRawHourlyTenDay(query *Query) ([]byte, error) {
	return w.Get(HourTenDay, query)
}

// GetHourlyTenDay returns the HourlyTenDay
func (w *Wug) GetHourlyTenDay(query *Query) (*HourlyTenDay, error) {
	data, err := w.GetRawHourlyTenDay(query)
	if err != nil {
		return nil, err
	}

	hourly := &HourlyTenDay{}
	err = json.Unmarshal(data, hourly)
	if err != nil {
		return nil, err
	}
	return hourly, nil
}

// GetRawForecast returns the raw bytes of a forecast request
func (w *Wug) GetRawForecast(query *Query) ([]byte, error) {
	return w.Get(Fore, query)
}

// GetForecast returns the Forecast
func (w *Wug) GetForecast(query *Query) (*Forecast, error) {
	data, err := w.GetRawForecast(query)
	if err != nil {
		return nil, err
	}

	forecast := &Forecast{}
	err = json.Unmarshal(data, forecast)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

// GetRawForecastTenDay returns the raw bytes of a ten day forecast request
func (w *Wug) GetRawForecastTenDay(query *Query) ([]byte, error) {
	return w.Get(ForeTenDay, query)
}

// GetForecastTenDay returns the ForecastTenDay
func (w *Wug) GetForecastTenDay(query *Query) (*ForecastTenDay, error) {
	data, err := w.GetRawForecastTenDay(query)
	if err != nil {
		return nil, err
	}

	forecast := &ForecastTenDay{}
	err = json.Unmarshal(data, forecast)
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

// Get returns the raw bytes of a request type given the provided Query.
func (w *Wug) Get(requestType RequestType, query *Query) ([]byte, error) {
	request := fmt.Sprintf(requestURL, query.apiKey, requestMap[requestType], query.queryValue)
	resp, err := w.Client.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
