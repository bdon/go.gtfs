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
	Shapes map[string]*Shape
}

type Route struct {
	Id        string
	ShortName string
	Trips     []Trip
}

type Trip struct {
	Id    string
	Shape *Shape
}

type Shape struct {
	Id     string
	Coords []Coord
}

type Coord struct {
	Lat float64
	Lon float64
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
	f.Shapes = make(map[string]*Shape)

	// we assume that this CSV is ordered by shape_id
	// but this is not guaranteed in spec?
	var curShape *Shape
	var found = false
	f.readCsv("shapes.txt", func(s []string) {
		shape_id := s[0]
		if !found || shape_id != curShape.Id {
			if found {
				f.Shapes[curShape.Id] = curShape
			}
			found = true
			curShape = &Shape{Id: shape_id}
		}
	})
	if found {
		f.Shapes[curShape.Id] = curShape
	}

	f.readCsv("routes.txt", func(s []string) {
		rsn := strings.TrimSpace(s[2])
		id := strings.TrimSpace(s[0])
		f.Routes[id] = Route{Id: id, ShortName: rsn}
	})

	f.readCsv("trips.txt", func(s []string) {
		route_id := s[0]
		shape_id := s[6]
		route := f.Routes[route_id]
		var shape *Shape
		shape = f.Shapes[shape_id]
		route.Trips = append(route.Trips, Trip{Shape: shape})
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
func (route Route) Shapes() []*Shape {
	// collect the unique list of shape pointers
	hsh := make(map[*Shape]bool)

	for _, v := range route.Trips {
		hsh[v.Shape] = true
	}

	retval := []*Shape{}
	for k, _ := range hsh {
		retval = append(retval, k)
	}
	return retval
}
