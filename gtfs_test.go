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
