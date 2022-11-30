package main

import (
	//"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"

	//"os"
	"strings"
)

func main() {
	//Sprintf is to print back string to another variable

	s := fmt.Sprintf("%[2]d %1[1]d \n", 40, 30)
	fmt.Print(s)

	f := 123.1212
	fmt.Printf("%.2f\n", f)

	//r := bufio.NewReader(os.Stdin)
	//s, _ = r.ReadString('\n')
	//fmt.Println("The input is:", s)

	//split string using fields function
	s = "This is a sooper long string, which has no stops."
	theFields := strings.Fields(s)
	fmt.Printf("%q\n", theFields)

	//function to return true if any of the character is punctuation
	func1 := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	result3 := strings.FieldsFunc(s, func1)
	fmt.Printf("%q\n", result3)

	//map function to map each character of string to other
	//string builder to append strings

	var sb strings.Builder
	sb.WriteString("this is string one \n")
	sb.WriteString("this is string two \n")
	sb.WriteString("this is string three \n")

	fmt.Println(sb.String())
	fmt.Println("capacity:", sb.Cap())

	sone, err := strconv.ParseInt("120", 10, 64)
	if err != nil {
		fmt.Println("its error:", err.Error())
	}
	fmt.Printf("%T %v\n", sone, sone)
	formatstring := strconv.FormatInt(98, 10)
	fmt.Printf("%T %v\n", formatstring, formatstring)

	//random numbers
	//initialise random seed to an unknown value
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))

	shufflestring := "one two three four five"
	words := strings.Fields(shufflestring)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
