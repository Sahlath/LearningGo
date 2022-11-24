package main

import (
	"fmt"
	"strings"
)

func main() {
	anInteger := 34
	var p = &anInteger
	fmt.Println("value of pointer p is::", *p)

	aFloat := 12.12
	pointer2 := &aFloat
	fmt.Println("value of pointer2 is::", *pointer2)

	*pointer2 = *pointer2 / 5
	fmt.Println("value of pointer2 now is::", *pointer2)
	fmt.Println("value of initial float is::", aFloat)

	text := "this is a long text"
	fmt.Println("the number of characters in the string is", len(text))
	fmt.Printf("value is %v: and type is %T", text, text)
	fmt.Println("first value is", string(text[0]))

	//multiline strings
	poem := `this is 
			a poem about sky
			sky is blue
			it is so cool
			`
	fmt.Println(poem)

	//words in the poem
	words := strings.Fields(poem)
	wordCount := map[string]int{}
	for _, value := range words {
		wordCount[value]++
	}

	fmt.Println("word count is", wordCount)
}
