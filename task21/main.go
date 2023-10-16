package main

import "fmt"

type StdPrinter struct{}

func (o *StdPrinter) Print(s string) {
    fmt.Println(s)
}

type Printer interface {
    PrintMessage(msg string)
}

type PrinterAdapter struct {
    StdPrinter *StdPrinter
}

func (pa *PrinterAdapter) PrintMessage(msg string) {
    pa.StdPrinter.Print(msg)
}

func main() {
  printer := &StdPrinter{}
  adapter := &PrinterAdapter{printer}
  adapter.PrintMessage("test")
}
