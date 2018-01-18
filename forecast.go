package wug

// ForecastDay returns a forecast for a day
type ForecastDay struct {
	Date struct {
		Epoch          string `json:"epoch"`
		Pretty         string `json:"pretty"`
		Day            int    `json:"day"`
		Month          int    `json:"month"`
		Year           int    `json:"year"`
		Yday           int    `json:"yday"`
		Hour           int    `json:"hour"`
		Min            string `json:"min"`
		Sec            int    `json:"sec"`
		Isdst          string `json:"isdst"`
		Monthname      string `json:"monthname"`
		MonthnameShort string `json:"monthname_short"`
		WeekdayShort   string `json:"weekday_short"`
		Weekday        string `json:"weekday"`
		Ampm           string `json:"ampm"`
		TzShort        string `json:"tz_short"`
		TzLong         string `json:"tz_long"`
	} `json:"date"`
	Period int `json:"period"`
	High   struct {
		Fahrenheit string `json:"fahrenheit"`
		Celsius    string `json:"celsius"`
	} `json:"high"`
	Low struct {
		Fahrenheit string `json:"fahrenheit"`
		Celsius    string `json:"celsius"`
	} `json:"low"`
	Conditions string `json:"conditions"`
	Icon       string `json:"icon"`
	IconURL    string `json:"icon_url"`
	Skyicon    string `json:"skyicon"`
	Pop        int    `json:"pop"`
	QpfAllday  struct {
		In float64 `json:"in"`
		Mm int     `json:"mm"`
	} `json:"qpf_allday"`
	QpfDay struct {
		In float64 `json:"in"`
		Mm int     `json:"mm"`
	} `json:"qpf_day"`
	QpfNight struct {
		In float64 `json:"in"`
		Mm int     `json:"mm"`
	} `json:"qpf_night"`
	SnowAllday struct {
		In float64 `json:"in"`
		Cm float64 `json:"cm"`
	} `json:"snow_allday"`
	SnowDay struct {
		In float64 `json:"in"`
		Cm float64 `json:"cm"`
	} `json:"snow_day"`
	SnowNight struct {
		In float64 `json:"in"`
		Cm float64 `json:"cm"`
	} `json:"snow_night"`
	Maxwind struct {
		Mph     int    `json:"mph"`
		Kph     int    `json:"kph"`
		Dir     string `json:"dir"`
		Degrees int    `json:"degrees"`
	} `json:"maxwind"`
	Avewind struct {
		Mph     int    `json:"mph"`
		Kph     int    `json:"kph"`
		Dir     string `json:"dir"`
		Degrees int    `json:"degrees"`
	} `json:"avewind"`
	Avehumidity int `json:"avehumidity"`
	Maxhumidity int `json:"maxhumidity"`
	Minhumidity int `json:"minhumidity"`
}

// TxtForecastDay text representation of a day of forecast information
type TxtForecastDay struct {
	Period        int    `json:"period"`
	Icon          string `json:"icon"`
	IconURL       string `json:"icon_url"`
	Title         string `json:"title"`
	Fcttext       string `json:"fcttext"`
	FcttextMetric string `json:"fcttext_metric"`
	Pop           string `json:"pop"`
}

// TxtForecast contains date and TxtForecastday
type TxtForecast struct {
	Date           string           `json:"date"`
	TxtForecastday []TxtForecastDay `json:"forecastday"`
}

// SimpleForecast contains the forecast information for N days
type SimpleForecast struct {
	Forecastday []ForecastDay `json:"forecastday"`
}

// ForecastData contains the Textforecast and Simpleforecast information.
type ForecastData struct {
	Textforecast   TxtForecast    `json:"txt_forecast"`
	Simpleforecast SimpleForecast `json:"simpleforecast"`
}

// Forecast current forecast
type Forecast struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Forecast int `json:"forecast"`
		} `json:"features"`
	} `json:"response"`
	Forecast ForecastData `json:"forecast"`
}

// ForecastTenDay the ten day forecast
type ForecastTenDay struct {
	Response struct {
		Version        string `json:"version"`
		TermsofService string `json:"termsofService"`
		Features       struct {
			Forecast10Day int `json:"forecast10day"`
		} `json:"features"`
	} `json:"response"`
	Forecast ForecastData `json:"forecast"`
}
