package main

import (
	"fmt"
)

// func main() {
// 	area, err := area(5, 4)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("area of rectange is:", *area)

// 	cities := [2]string{""} //array
// 	copy := cities
// 	cities[1] = "Bangalore"
// 	fmt.Println("copy is", copy)
// 	addCity("Miami", copy)
// 	addCityByReference("Miami", &copy)
// 	fmt.Println("cities is", cities)
// 	fmt.Println("copy is", copy)

// 	mycities := make([]string, 2, 2)
// 	mycopy := mycities
// 	mycities[0] = "Kannur"
// 	fmt.Println("mycopy is", mycopy)

// 	//working with maps
// 	m := map[string]float64{
// 		"Chennai": 40.4,
// 		"Kannur":  30,
// 	}
// 	m["Bangalore"] = 24.7

// 	//to find the hottest city
// 	city, temp := hottestCity(m)
// 	fmt.Println("city is", city)
// 	fmt.Println("temp is", temp)
// }

func area(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero values %v %v", x, y)
	}
	area := x * y
	return &area, nil
}

//arrays are value types, not reference types,
//if we pass arrays to a method and update there, it will not change the actual value

func addCity(city string, cities [2]string) {
	cities[1] = city
}

func addCityByReference(city string, cities *[2]string) {
	cities[1] = city
}

func hottestCity(cities map[string]float64) (string, float64) {
	var hottestCity string
	var maxTemp float64
	for c, v := range cities {
		if v > maxTemp {
			maxTemp = v
			hottestCity = c
		}
	}
	return hottestCity, maxTemp

}
