package gtfs

import (
	"encoding/csv"
	"io"
	"os"
	"path"
	"sort"
	"strconv"
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
	Seq int
}

type BySeq []Coord

func (a BySeq) Len() int           { return len(a) }
func (a BySeq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySeq) Less(i, j int) bool { return a[i].Seq < a[j].Seq }

// main utility function for reading GTFS files
func (feed *Feed) readCsv(filename string, f func([]string)) {
	file, err := os.Open(path.Join(feed.Dir, filename))
	if err != nil {
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

	// we assume that this CSV is grouped by shape_id
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
		lon, _ := strconv.ParseFloat(s[1], 64)
		lat, _ := strconv.ParseFloat(s[2], 64)
		seq, _ := strconv.Atoi(s[3])
		curShape.Coords = append(curShape.Coords, Coord{Lat: lat, Lon: lon, Seq: seq})
	})
	if found {
		f.Shapes[curShape.Id] = curShape
	}

	// sort coords by their sequence
	for _, v := range f.Shapes {
		sort.Sort(BySeq(v.Coords))
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
