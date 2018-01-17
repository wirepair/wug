package wug

import (
	"testing"
)

func TestQueryTypes(t *testing.T) {
	q := NewQueryByAutoIP("apikey")
	if q.queryType != AutoIP {
		t.Fatalf("expected AutoIP")
	}

	q = NewQueryByAirportCode("apikey", "NRT")
	if q.queryType != AirportCode {
		t.Fatalf("expected AirportCode")
	}

	q = NewQueryByCountryCity("apikey", "jp", "tokyo")
	if q.queryType != CountryCity {
		t.Fatalf("expected CountryCity")
	}

	q = NewQueryByIPGeo("apikey", "127.0.0.1")
	if q.queryType != IPGeo {
		t.Fatalf("expected IPGeo")
	}

	q = NewQueryByLatLong("apikey", "37.8", "-122.4")
	if q.queryType != LatLong {
		t.Fatalf("expected LatLong")
	}

	q = NewQueryByPwsID("apikey", "KCASANFR70")
	if q.queryType != PwsID {
		t.Fatalf("expected PwsID")
	}

	q = NewQueryByUsStateCity("apikey", "CA", "san francisco")
	if q.queryType != UsStateCity {
		t.Fatalf("expected UsStateCity")
	}

	q = NewQueryByUsZip("apikey", "90210")
	if q.queryType != UsZip {
		t.Fatalf("expected UsZip")
	}
}
