package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("Current time is:", n)

	joiningDate := time.Date(2022, time.November, 8, 11, 0, 0, 0, time.UTC)
	fmt.Println("Joining date is \n", joiningDate)
	fmt.Printf("Joining date is %T\n", joiningDate)
	fmt.Println("Formatted date is \n", joiningDate.Format(time.ANSIC))

}
