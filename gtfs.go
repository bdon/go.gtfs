package gtfs

import (
	"encoding/csv"
	"io"
	"os"
	"path"
	"strings"
)

type Feed struct {
	Dir    string
	Routes map[string]Route
}

type Route struct {
	Id        string
	ShortName string
	Trips     []Trip
}

type Trip struct {
	Id string
}

type Shape struct {
}

// main utility function for reading GTFS files
func (feed *Feed) readCsv(filename string, f func([]string)) {
	file, err := os.Open(path.Join(feed.Dir, filename))
	if err != nil {
		p(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.TrailingComma = true
	firstLineSeen := false
	for {
		record, err := reader.Read()
		if !firstLineSeen {
			firstLineSeen = true
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
		} else {
			f(record)
		}
	}
}

func Load(feed_path string) Feed {
	f := Feed{Dir: feed_path}
	f.Routes = make(map[string]Route)

	f.readCsv("routes.txt", func(s []string) {
		rsn := strings.TrimSpace(s[2])
		id := strings.TrimSpace(s[0])
		f.Routes[id] = Route{Id: id, ShortName: rsn}
	})

	f.readCsv("trips.txt", func(s []string) {
		route_id := s[0]
		route := f.Routes[route_id]
		route.Trips = append(route.Trips, Trip{Id: "5"})
		f.Routes[route_id] = route
	})

	return f
}

func (feed *Feed) RouteByShortName(shortName string) Route {
	for _, v := range feed.Routes {
		if v.ShortName == shortName {
			return v
		}
	}
	//TODO error here
	return Route{}
}

// get All shapes for a route
//
func (route *Route) Shapes() []Shape {
	return []Shape{}
}
