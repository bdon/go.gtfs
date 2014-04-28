package gtfs

import (
	"testing"
)

func TestLoadRoutes(t *testing.T) {
	feed := Load("./fixtures")

	if len(feed.Routes) < 1 {
		t.Error("Feed should have at least one route")
	}

	res := feed.Routes["1093"].Id
	if res != "1093" {
		t.Errorf("Feed should have route IDs, got %s", res)
	}

	res = feed.RouteByShortName("N").ShortName
	if res != "N" {
		t.Errorf("Feed should be addressable by Route short name, got %s", res)
	}
}

func TestTrips(t *testing.T) {
	feed := Load("./fixtures")

	res := len(feed.RouteByShortName("N").Trips)
	if res != 23 {
		t.Errorf("Route should have all trips, got %d", res)
	}
}

func TestShapes(t *testing.T) {
	feed := Load("./fixtures")
	res := len(feed.Shapes)
	if res != 3 {
		t.Errorf("Feed should have three total shapes")
	}

	res = len(feed.RouteByShortName("N").Shapes())
	if res != 2 {
		t.Errorf("Route should have two total shapes")
	}

	coord := feed.Shapes["116466"].Coords[0]
	if !(coord.Lat == 37.760036 && coord.Lon == -122.509075) {
		t.Errorf("Failed to parse shape coordinates.")
	}
}

func TestTimeParsing(t *testing.T) {
	res := Hmstoi("00:00:00")
	expected := 0
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
	res = Hmstoi("23:59:59")
	expected = 86399
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
	res = Hmstoi("12:34:56")
	expected = 45296
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestStops(t *testing.T) {
	feed := Load("./fixtures")
	res := len(feed.Stops)
	if res != 4 {
		t.Errorf("Feed should have four total stops")
	}
}
