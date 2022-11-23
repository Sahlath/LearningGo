package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 20, 30, 20
	sum := i1 + i2 + i3
	fmt.Println("Sum is :", sum)

	f1, f2, f3 := 23.7, 24.5, 12.9
	sumFloat := f1 + f2 + f3
	fmt.Println("sumFloat is :", sumFloat)

	sumFloat = math.Round(sumFloat)
	fmt.Println("sumFloat rounded is :", sumFloat)

	circleRadius := 12.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumefernce is %.3f\n", circumference)
}
