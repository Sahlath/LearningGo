package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func main() {
	fmt.Println("GOOOOO!!!")
	newBudget := Budget{"xyz", 12.34, time.Now().Add(time.Duration(10))}
	fmt.Println("TimeLeft", newBudget.getTimeLeft())

	//create square and calculate its area

	sq1, err := NewSquare(1, 1, 10)
	if err != nil {
		fmt.Println("Cant create square")
		log.Fatalf("Cant create square")
	}
	sq1.Move(3, 5)
	fmt.Printf("Square is %+v\n", sq1)
	fmt.Println("Area is", sq1.Area())

	c := Circle{3.5}
	shapes := []Shape{c, sq1}
	fmt.Println(sumOfAreas(shapes))

	//get type of urls
	urls := []string{"https://golang.org", "https://linkedin.org", "https://api.github.com"}
	ch1, ch2 := make(chan string), make(chan int)
	for _, url := range urls {
		go returnType(url, ch1)
	}

	for range urls {
		select {
		case out := <-ch1:
			fmt.Println("got from channel1", out)
		case out := <-ch2:
			fmt.Println("got from channel2", out)
		}
	}

	//use context with time out to find the best bid
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	url := "http://abc.com"
	bid := findBid(ctx, url)
	fmt.Println(bid)
}

type Budget struct {
	CampaignID  string
	Balance     float64
	timeExpires time.Time
}

func (b *Budget) getTimeLeft() int {
	//return int(b.timeExpires.Sub(time.Now()))
	return int(time.Until(b.timeExpires))
}

type Square struct {
	X      int
	Y      int
	Length float64
}

func NewSquare(x int, y int, length float64) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("cannot create square")
	}

	s := Square{
		X:      x,
		Y:      y,
		Length: length,
	}

	return &s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

// Interfaces..if our structs implement this method, all of them can be considered as sma
type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * 2 * c.radius * c.radius
}

func sumOfAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// using channels and routines
func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> ctype: %s", url, ctype)
}

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)

	return Bid{
		AdURL: url,
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "bluble.com",
	Price: 23.12,
}

// Use of contexts
func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered to avoid goroutine leak
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}
