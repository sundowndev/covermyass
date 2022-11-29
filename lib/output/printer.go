package output

import (
	"fmt"
	"os"
)

type Printer interface {
	Printf(string, ...interface{})
	Println(string)
}

var globalPrinter Printer = &VoidPrinter{}

func ChangePrinter(printer Printer) {
	globalPrinter = printer
}

func Printf(format string, args ...interface{}) {
	globalPrinter.Printf(format, args...)
}

func Println(format string) {
	globalPrinter.Printf(format)
}

type ConsolePrinter struct{}

func NewConsolePrinter() Printer {
	return &ConsolePrinter{}
}

func (c *ConsolePrinter) Printf(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
}

func (c *ConsolePrinter) Println(format string) {
	_, _ = fmt.Fprintln(os.Stdout, format)
}

type VoidPrinter struct{}

func (v *VoidPrinter) Printf(_ string, _ ...interface{}) {}

func (v *VoidPrinter) Println(_ string) {}
