package models

import (
	"fmt"
	"javatransition/samples/javatransition/data"
	"sort"
)

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	ListAll() []CityTemp
	Filter(beach bool, ski bool) []CityTemp
}

// NewCities initialises the Cities data structure by calling the
// ReadData method to read information from file.
func NewCities(reader data.DataReader) (Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp)
	for _, r := range d {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountain)
	}

	return &cities{
		cityMap: cmap,
	}, nil
}

// ListAll returns a slice of all the cities.
func (c cities) ListAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	sortAlphabetically(cs)
	return cs
}

func sortAlphabetically(cs []CityTemp) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() > cs[j].Name()
	})
}

// Filter process the beach and ski markers and returns the list cities.
func (c cities) Filter(beach bool, ski bool) []CityTemp {
	// no flags set, return all
	fmt.Println("beach is", beach)
	fmt.Println("ski is", ski)
	if !beach && !ski {
		return c.ListAll()
	}
	return c.filterHelper(beach, ski)
}

// filterHelper is a helper method that processes the values of
// beach and ski flags on the list of cities.
func (c cities) filterHelper(beach bool, ski bool) []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		if matchFilter(rc, beach, ski) {
			fmt.Println("if match is true")
			cs = append(cs, rc)
		}
	}
	fmt.Println("length of cs after filter", len(cs))
	sortAlphabetically(cs)
	return cs
}

// matchFilter returns whether the given city matches given filter parameters
func matchFilter(rc CityTemp, beach bool, ski bool) bool {
	if beach && rc.BeachVacationReady() {
		return true
	}
	if ski && rc.SkiVacationReady() {
		return true
	}

	return false
}
