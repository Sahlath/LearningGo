package printer

import (
	"fmt"
	"javatransition/samples/javatransition/models"
	"os"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
}

// New initialises a new Printer instance
func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

// CityHeader prints out the table header for the city table
func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach ready?\tSki ready?")
}

// CityDetails prints out the given city
func (p *Printer) CityDetails(c models.CityTemp) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name(), c.TempC(), c.TempF(), c.BeachVacationReady(), c.SkiVacationReady())
}

// CityDetails prints out the given city
func (p *Printer) CityDetails1(c models.CityTemp1, cq models.CityQuery) {
	fmt.Fprintf(p.w, "%v\t%v\t%v\t%v\t%v\n", c.Name1(), c.TempC1(cq), c.TempF1(cq), c.BeachVacationReady1(cq), c.SkiVacationReady1(cq))
}

// Cleanup closes up the printer instance
func (p *Printer) Cleanup() {
	p.w.Flush()
}
