package output

import (
	"fmt"
	"os"
)

type Printer interface {
	Printf(string, ...interface{})
	Println(string)
	Errorf(string, ...interface{}) error
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

func Errorf(format string, args ...interface{}) error {
	return globalPrinter.Errorf(format, args...)
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

func (c *ConsolePrinter) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

type VoidPrinter struct{}

func (v *VoidPrinter) Printf(_ string, _ ...interface{}) {}

func (v *VoidPrinter) Println(_ string) {}

func (v *VoidPrinter) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}
