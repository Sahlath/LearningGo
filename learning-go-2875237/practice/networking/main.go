package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// pkg.go.dev/
func main() {
	rdr := strings.NewReader(data)
	dec := json.NewDecoder(rdr)
	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("Cant decode = %s\n", err)
	}

	fmt.Printf("got : %+v\n", req)

	//Create response
	resp := map[string]interface{}{
		"ok":      true,
		"balance": req.Amount,
	}

	//Encode the response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("cant encode - %s", err)
	}

	//HTTP post calls

	req1 := &Request{
		Login:  "Sahla",
		Type:   "employee",
		Amount: 12000,
	}

	var buf bytes.Buffer
	enc1 := json.NewEncoder(&buf)
	if err := enc1.Encode(req1); err != nil {
		log.Fatalf("cant encode - %s", err)
	}

	resp1, err1 := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err1 != nil {
		log.Fatalf("cant call url")
	}

	defer resp1.Body.Close()

	io.Copy(os.Stdout, resp1.Body)

	//setting timeout to http call and limit to the size of data received
	/*ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	req2, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp2, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Fatal(err)
	}
	defer resp2.Body.Close()

	const mb = 200
	r := io.LimitReader(resp2.Body, mb)
	io.Copy(os.Stdout, r)*/

	//call git hub and my public repos
	/*user, err := userInfo("Sahlath")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)*/

	//working with HTTP server
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	log.Printf("Server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

var data = `
{
	"user":"Sahla",
	"type":"employee",
	"amount":12000
}`

func userInfo(user string) (*User, error) {
	//HTTP call
	call := fmt.Sprintf("https://api.github.com/users/%s", url.PathEscape(user))
	resp, err := http.Get(call)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//Decode the response JSON
	userInfo := User{Login: user}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&userInfo); err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// healthHandler returns server health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	//Step 1: Decode & validate the request
	defer r.Body.Close()
	req := &MathRequest{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(req); err != nil {
		log.Printf("Bad json : %s", err)
		http.Error(w, "Bad json", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("Bad operator : %q", req.Op)
		http.Error(w, "Unknown Operator", http.StatusBadRequest)
		return
	}

	//2. Do the work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "Divide by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("Unknown operation :%s", req.Op)
	}

	//3. Encode the result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("cant encode %v - %s", resp, err)
	}
}
