package models

var (
	// hot enough for the beach
	beachVacationThreshold float64 = 22
	// cold enough to snow
	skiVacationThreshold float64 = -2
)

type city struct {
	name        string
	tempC       float64 // temperature in Celsius
	hasBeach    bool
	hasMountain bool
}

type city1 struct {
	name        string
	tempC       []float64 // temperature in Celsius
	hasBeach    bool
	hasMountain bool
}

// CityTemp interface defines all things city temperature related
type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

// CityTemp interface defines all things city temperature related
type CityTemp1 interface {
	Name1() string
	TempC1(q CityQuery) float64
	TempF1(q CityQuery) float64
	BeachVacationReady1(q CityQuery) bool
	SkiVacationReady1(q CityQuery) bool
}

// NewCity creates a new city instance with the given name and Celsius temperated
func NewCity(n string, t float64, hasBeach bool, hasMountain bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

// NewCity creates a new city instance with the given name and Celsius temperated
func NewCity1(n string, t []float64, hasBeach bool, hasMountain bool) CityTemp1 {
	return &city1{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

// Name returns the city name
func (c city) Name() string {
	return c.name
}

// TempC returns the city temperature in Celsius
func (c city) TempC() float64 {
	return c.tempC
}

// TempF returns the city temperature converted to Fahrenheit
func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

// BeachVacationReady returns whether the given city is suitable for a beach vacation.
func (c city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}

// SkiVacationReady returns whether the given city is suitable for a ski vacation.
func (c city) SkiVacationReady() bool {
	return c.hasMountain && c.tempC < skiVacationThreshold
}

//for the city1

// Name returns the city name
func (c city1) Name1() string {
	return c.name
}

// TempC returns the city temperature in Celsius
func (c city1) TempC1(q CityQuery) float64 {
	return c.tempC[q.Month()-1]
}

// TempF returns the city temperature converted to Fahrenheit
func (c city1) TempF1(q CityQuery) float64 {
	return (c.tempC[q.Month()-1] * 9 / 5) + 32
}

// BeachVacationReady returns whether the given city is suitable for a beach vacation.
func (c city1) BeachVacationReady1(q CityQuery) bool {
	return c.hasBeach && c.tempC[q.Month()-1] > beachVacationThreshold
}

// SkiVacationReady returns whether the given city is suitable for a ski vacation.
func (c city1) SkiVacationReady1(q CityQuery) bool {
	return c.hasMountain && c.tempC[q.Month()-1] < skiVacationThreshold
}
