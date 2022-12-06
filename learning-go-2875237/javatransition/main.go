package main

import (
	"flag"
	"fmt"
	"javatransition/samples/javatransition/data"
	"javatransition/samples/javatransition/models"
	"javatransition/samples/javatransition/printer"
)

func main() {
	fmt.Printf("Welcome to the LinkedIn Learning Temperature Service!\n\n")

	//Use of flags to make CLI
	beachReady := flag.Bool("beach", false, "display only beach ready ones")
	skiReady := flag.Bool("ski", false, "display only ski ready ones")
	month := flag.Int("month", 0, "look up for destinations in a given month")
	name := flag.String("name", "", "look up for destinations by name")
	flag.Parse()

	//create query with flags
	q, err := models.NewQuery(*beachReady, *skiReady, *month, *name)
	fmt.Println("Query is: ", q)
	if err != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}
	//cities, err := models.NewCities(data.NewReader1())
	cities1, err1 := models.NewCities1(data.NewReader1())
	if err != nil || err1 != nil {
		fmt.Println("Fatal error occurred: ", err)
		return
	}

	// initialise printer and defer cleanup
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	//filter cities
	//cs := cities.Filter(*beachReady, *skiReady)
	cs1 := cities1.Filter1(q)
	//print all the cities
	// for _, c := range cs {
	// 	p.CityDetails(c)
	// }
	for _, c := range cs1 {
		p.CityDetails1(c, q)
	}

	//print all the cities
	// for _, c := range cities.ListAll() {
	// 	p.CityDetails(c)
	// }
}
