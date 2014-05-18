package gtfs

import (
	"testing"
)

func TestLoadRoutes(t *testing.T) {
	feed := Load("./fixtures", true)

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

	res = feed.RouteByShortName("N").LongName
	if res != "JUDAH" {
		t.Errorf("Feed should be have long name, got %s", res)
	}
}

func TestTrips(t *testing.T) {
	feed := Load("./fixtures", true)

	res := len(feed.Trips)
	if res != 24 {
		t.Errorf("Feed should have 24 trips")
	}

	res = len(feed.RouteByShortName("N").Trips)
	if res != 23 {
		t.Errorf("Route should have all trips, got %d", res)
	}

	firstTrip := feed.RouteByShortName("N").Trips[0]
	if firstTrip.Service != "1" {
		t.Errorf("trip should have Service")
	}

	if firstTrip.Direction != "1" {
		t.Errorf("trip should have Direction")
	}
}

func TestStopTimes(t *testing.T) {
	feed := Load("./fixtures", true)
	trip := feed.RouteByShortName("N").Trips[0]
	res := len(trip.StopTimes)
	if res != 2 {
		t.Errorf("Trip should have two stop times.")
	}

	stopTime1 := trip.StopTimes[0]

	res = stopTime1.Seq
	if res != 1 {
		t.Errorf("StopTime should have seq 1, got %d", res)
	}

	res = stopTime1.Time
	if res != 25740 {
		t.Errorf("Stop with sequence 1 should have time 25740, got %d", res)
	}
}

func TestShapes(t *testing.T) {
	feed := Load("./fixtures", true)
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

func TestLongestShape(t *testing.T) {
	// this is kind of a hack.
	// But for convenience, determine the longest shape of a Route.
	// By # of stops perhaps? ensure that it is a superset of other trips?

	feed := Load("./fixtures", true)
	res := feed.RouteByShortName("N").LongestShape()
	if res.Id != "116466" {
		t.Errorf("Longest shape should be 116466, got %s", res.Id)
	}
}

func TestRouteStops(t *testing.T) {
	// this won't work for forking routes

	feed := Load("./fixtures", true)
	res := len(feed.RouteByShortName("N").Stops())
	if res != 2 {
		t.Errorf("Route should have 2 total stops, got %d", res)
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
	feed := Load("./fixtures", true)
	res := len(feed.Stops)
	if res != 4 {
		t.Errorf("Feed should have four total stops")
	}
}
