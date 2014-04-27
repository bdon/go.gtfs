package gtfs

import (
	"fmt"
	"testing"
)

func p(s interface{}) {
	fmt.Println(s)
}

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

}
