package wug

import (
	"fmt"
	"strings"
)

// QueryType of query to build
type QueryType int

// The query type definitions (how to search)
const (
	PwsID       QueryType = iota // by PWS station id
	UsStateCity                  // by us state and city
	UsZip                        // by us zip code
	CountryCity                  // by country / city
	AirportCode                  // by airport code
	LatLong                      // by latitude and longitude
	AutoIP                       // by automatic geo ip address
	IPGeo                        // by provided ip address
)

var queryFormats = map[QueryType]string{
	PwsID:       "/pws:%s.json",
	UsStateCity: "/%s/%s.json",
	UsZip:       "/%s.json",
	CountryCity: "/%s/%s.json",
	LatLong:     "/%s,%s.json",
	AirportCode: "/%s.json",
	AutoIP:      "/autoip.json",
	IPGeo:       "/autoip.json?geo_ip=%s",
}

// Query used for the Wug client
type Query struct {
	apiKey     string
	queryType  QueryType
	queryValue string
}

// NewQueryByPwsID query by pws id, pwsID does not need the leading pws: string
func NewQueryByPwsID(apiKey string, pwsID string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  PwsID,
		queryValue: fmt.Sprintf(queryFormats[PwsID], pwsID),
	}
}

// NewQueryByUsStateCity query by US state and city, replaces spaces with _ and
// upper cases the state.
func NewQueryByUsStateCity(apiKey string, state, city string) *Query {
	state = strings.Replace(state, " ", "_", -1)
	state = strings.ToUpper(state)
	city = strings.Replace(city, " ", "_", -1)

	return &Query{
		apiKey:     apiKey,
		queryType:  UsStateCity,
		queryValue: fmt.Sprintf(queryFormats[UsStateCity], state, city),
	}
}

// NewQueryByUsZip query by US zip code
func NewQueryByUsZip(apiKey string, zip string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  UsZip,
		queryValue: fmt.Sprintf(queryFormats[UsZip], zip),
	}
}

// NewQueryByCountryCity query by country and city.
func NewQueryByCountryCity(apiKey string, country, city string) *Query {
	country = strings.Replace(country, " ", "_", -1)
	city = strings.Replace(city, " ", "_", -1)

	return &Query{
		apiKey:     apiKey,
		queryType:  CountryCity,
		queryValue: fmt.Sprintf(queryFormats[CountryCity], country, city),
	}
}

// NewQueryByLatLong query by latitude and longitude.
func NewQueryByLatLong(apiKey string, latitude, longitude string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  LatLong,
		queryValue: fmt.Sprintf(queryFormats[LatLong], latitude, longitude),
	}
}

// NewQueryByAirportCode query by airport code
func NewQueryByAirportCode(apiKey string, airport string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  AirportCode,
		queryValue: fmt.Sprintf(queryFormats[AirportCode], airport),
	}
}

// NewQueryByAutoIP query by geolocating the requester IP address and using
// the closest station.
func NewQueryByAutoIP(apiKey string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  AutoIP,
		queryValue: queryFormats[AutoIP],
	}
}

// NewQueryByIPGeo query by geolocating the provided IP address and using
// the closest station.
func NewQueryByIPGeo(apiKey string, ipAddress string) *Query {
	return &Query{
		apiKey:     apiKey,
		queryType:  IPGeo,
		queryValue: fmt.Sprintf(queryFormats[IPGeo], ipAddress),
	}
}

// Format the requestURL for the query with the query value.
func (q *Query) Format(requestURL string) string {
	return fmt.Sprintf(requestURL, q.queryValue)
}
