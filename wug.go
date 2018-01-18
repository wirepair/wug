package wug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const requestUrl = "http://api.wunderground.com/api/%s/%s/q/%s"

type RequestType int

const (
	Cond RequestType = iota
	Fore
	ForeTenDay
	Hour
	HourTenDay
)

var requestMap = map[RequestType]string{
	Cond:       "conditions",
	Fore:       "forecast",
	ForeTenDay: "forecast10day",
	Hour:       "hourly",
	HourTenDay: "hourly10day",
}

type Wug struct {
	client *http.Client
}

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
	return &Wug{client: client}
}

func (w *Wug) SetClient(client *http.Client) {
	w.client = client
}

func (w *Wug) GetRawConditions(query *Query) ([]byte, error) {
	return w.Get(Cond, query)
}

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

func (w *Wug) GetRawHourly(query *Query) ([]byte, error) {
	return w.Get(Hour, query)
}

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

func (w *Wug) GetRawHourlyTenDay(query *Query) ([]byte, error) {
	return w.Get(HourTenDay, query)
}

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

func (w *Wug) GetRawForecast(query *Query) ([]byte, error) {
	return w.Get(Fore, query)
}

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

func (w *Wug) GetRawForecastTenDay(query *Query) ([]byte, error) {
	return w.Get(ForeTenDay, query)
}

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

func (w *Wug) Get(requestType RequestType, query *Query) ([]byte, error) {
	request := fmt.Sprintf(requestUrl, query.apiKey, requestMap[requestType], query.queryValue)
	resp, err := w.client.Get(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
