package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

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

	//double the content of slice at a particular position
	//slice passes pointers to the values, hence change done in function reflects
	values := []int{3, 5, 6, 7}
	doubleAt(values, 3)
	fmt.Println(values)

	//pass by refernce
	val := 2
	double(val)
	fmt.Println(val)

	//pass by pointers
	doublePtr(&val)
	fmt.Println(val)

	ctype, err := getContentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("Error %v", err)
	} else {
		fmt.Println("Content type is", ctype)
	}
}

func doubleAt(values []int, at int) {
	values[at] *= 2
}

func double(val int) {
	val *= 2
}

func doublePtr(val *int) {
	*val *= 2
}

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("cannot find content type in header")
	}
	return contentType, nil
}
