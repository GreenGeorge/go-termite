package termite

import (
	"fmt"

	"github.com/buger/goterm"
	humanize "github.com/dustin/go-humanize"
)

func format(v int64, unit string) string {
	return fmt.Sprintf("%16s %s", humanize.Comma(v), unit)
}

type Output struct {
	TermWidth   int
	TableFormat string
}

func (o *Output) Rule() {
	for i := 0; i < o.TermWidth; i++ {
		goterm.Printf("%s", "-")
	}
	goterm.Printf("\n")
}

func (o *Output) PrintRow(description string, value int64, unit string) {
	goterm.Printf(o.TableFormat, description, format(value, unit))
}

func (o *Output) PrintHeader(header string) {
	o.PrintSpace(2)
	goterm.Println(header)
	o.Rule()
}

func (o *Output) PrintSpace(amount int) {
	for i := 0; i < amount; i++ {
		goterm.Printf("\n")
	}
}
