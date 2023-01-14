package hello

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Print() {
	fmt.Fprintf(p.Output, "Hello World")
}

func Print() {
	NewPrinter().Print()
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}
