package main

import (
	"fmt"
)

func main() {
	var aString string = "This is GO!"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString)
	var anotherString = "I am another String"
	fmt.Println(anotherString)
	fmt.Printf("The variable type is %T\n", anotherString)

	var anInteger int = 40
	fmt.Println(anInteger)
	fmt.Printf("The variable type is %T\n", anInteger)

	var anotherInteger float64 = 40.0
	fmt.Println(anotherInteger)
	fmt.Printf("The variable type is %T\n", anotherInteger)

	alsoString := "I am also a string"
	fmt.Println(alsoString)
	fmt.Printf("The variable type is %T\n", alsoString)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The defaultInt variable type is %T\n", defaultInt)

}
