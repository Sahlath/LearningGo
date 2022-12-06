package models

import (
	"fmt"
	"javatransition/samples/javatransition/data"
	"sort"
	"strings"
)

type cities1 struct {
	cityMap map[string]CityTemp1
}

type Cities1 interface {
	ListAll1() []CityTemp1
	Filter1(q CityQuery) []CityTemp1
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities1(reader data.DataReader) (Cities1, error) {
	d, err := reader.ReadData1()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp1)
	for _, r := range d {
		cmap[r.Id] = NewCity1(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &cities1{
		cityMap: cmap,
	}, nil
}

// ListAll returns a slice of all the cities.
func (c cities1) ListAll1() []CityTemp1 {
	var cs []CityTemp1
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	sortAlphabetically1(cs)
	return cs
}

func sortAlphabetically1(cs []CityTemp1) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name1() > cs[j].Name1()
	})
}

// Filter process the beach and ski markers and returns the list cities.
func (c cities1) Filter1(q CityQuery) []CityTemp1 {
	// no flags set, return all
	if !q.Beach() && !q.Ski() && q.Name() == "" {
		return c.ListAll1()
	}
	return c.filterHelper1(q)
}

// filterHelper is a helper method that processes the values of
// beach and ski flags on the list of cities.
func (c cities1) filterHelper1(q CityQuery) []CityTemp1 {
	var cs []CityTemp1
	for _, rc := range c.cityMap {
		if matchFilter1(rc, q) {
			fmt.Println("if match is true")
			cs = append(cs, rc)
		}
	}
	fmt.Println("length of cs after filter", len(cs))
	sortAlphabetically1(cs)
	return cs
}

// matchFilter returns whether the given city matches given filter parameters
func matchFilter1(rc CityTemp1, q CityQuery) bool {
	if q.Beach() && rc.BeachVacationReady1(q) {
		return true
	}
	if q.Ski() && rc.SkiVacationReady1(q) {
		return true
	}
	if q.Name() != "" && strings.Contains(strings.ToLower(rc.Name1()), q.Name()) {
		return true
	}

	return false
}
