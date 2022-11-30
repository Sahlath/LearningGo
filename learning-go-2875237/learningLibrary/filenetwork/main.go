package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func checkFileExists(filepath string) bool {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	fmt.Println("file name", fileinfo.Name())
	fmt.Println("file name", fileinfo.Mode())
	return true
}

func main() {
	exists := checkFileExists("sample.txt")
	fmt.Println("file exists??", exists)

	// data1 := []byte("this is some data")
	// ioutil.WriteFile("datafile.txt", data1, 0666)

	//another way to write file
	// f, _ := os.Create("datafile.txt")

	// defer f.Close()
	// f.WriteString("this is another set of string")
	// data2 := []byte{'s', 'a', 'h', 'l', 'a'}
	// f.Write(data2)

	// f.Sync()

	//truncate file
	// if checkFileExists("datafile.txt") {
	// 	os.Truncate("datafile.txt", 10)
	// }
	//append data to file
	//appendData("datafile.txt", "\nappending more")
	//readFile()

	//working with directories
	os.Mkdir("newdir", os.ModePerm)

	os.MkdirAll("com/kantar/newdir", os.ModePerm)

	//os.Remove("newdir")
	//os.RemoveAll("com")
	//loop through direcotries
	// dircontents, _ := ioutil.ReadDir(".")
	// for _, fi := range dircontents {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	// //working with urls
	// s := "https://www.example.com/user?username=sahla"

	// result, _ := url.Parse(s)
	// fmt.Println(result.RawQuery)
	// q := result.Query()
	// fmt.Println(q["username"])

	// newurl := url.URL{
	// 	Scheme:   "https",
	// 	Host:     "www.example.com",
	// 	Path:     "/args",
	// 	RawQuery: "x=1&y=2",
	// }
	// uu := newurl.String()
	// fmt.Println(uu)

	// //create new values struct and encode as query string
	// newvals := url.Values{}
	// newvals.Add("x", "100")
	// newvals.Add("z", "somestring")
	// newurl.RawQuery = newvals.Encode()

	// uu = newurl.String()
	// fmt.Println(uu)

	//GET operation
	const httpbin = "https://httpbin.org/get"
	//perform GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("status code::", resp.StatusCode)
	fmt.Println("status::", resp.Status)
	fmt.Println("content length::", resp.ContentLength)

	var sb strings.Builder
	contents, _ := ioutil.ReadAll(resp.Body)
	bufferLength, _ := sb.Write(contents)
	fmt.Println(bufferLength, sb.String())

	//
	//encodeExample()
	//decodeExample()
	//encodeXMLExample()
	decodeXMLExample()
}

func appendData(fname string, data string) {
	f, _ := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	_, err := f.WriteString(data)
	fmt.Println(err)

}

func readFile() {
	f, _ := os.Open("sample.txt")
	defer f.Close()
	b1 := make([]byte, 20)

	for {
		n, err := f.Read(b1)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Println("Bytes read::", n)
		fmt.Println("content is::", string(b1[:n]))

	}
}

type person struct {
	Name    string   `json:"name"`
	Address string   `json:"addr"`
	Age     int      `json:"age"`
	Colours []string `json:"favcolors,omitempty"`
}

type personxml struct {
	XMLName xml.Name `xml:"persondata"`
	Name    string   `xml:"name"`
	Address string   `xml:"addr"`
	Age     int      `xml:"age,attr"`
	Colours []string `xml:"colors"`
}

func encodeExample() {
	people := []person{
		{"Sahla", "Bangalore", 38, nil},
		{"Navas", "Bangalore", 44, []string{"blue", "white"}},
	}

	//encode array to JSON
	result, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)

}

func decodeExample() {
	//declare some json data to decode
	data := []byte(`{
	"name": "fannaan",
	"addr"	:"bangalore",
	"age"	:15,
	"favcolors"	: ["Orange","White"]
	}`)

	var p person
	valid := json.Valid(data)
	fmt.Println("isValid", valid)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

	//if json is not in correct structire
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Printf("key (%v), Value (%T %v)\n", k, v, v)
	}

}

func encodeXMLExample() {
	people := []personxml{
		{Name: "Sahla", Address: "Bangalore", Age: 38, Colours: nil},
		{Name: "Navas", Address: "Bangalore", Age: 44, Colours: []string{"blue", "white"}},
	}

	//encode array to JSON
	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%s\n", xml.Header, result)

}

func decodeXMLExample() {
	//declare some json data to decode
	xmldata := ` <persondata age="44">
        <name>Navas</name>
        <addr>Bangalore</addr>
        <colors>blue</colors>
        <colors>white</colors>
</persondata>
	`

	var p personxml

	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)

}
