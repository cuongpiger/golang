package main

import "fmt"

type (
	// Computer is the abstraction
	Computer interface {
		Print()
		SetPrinter(printer Printer)
	}

	// Printer is the implementation
	Printer interface {
		PrintFile()
	}

	// Mac is the refined abstraction
	Mac struct {
		printer Printer
	}

	// Windows is the refined abstraction
	Windows struct {
		printer Printer
	}

	// Epson is the concrete implementation
	Epson struct{}

	// HP is the concrete implementation
	HP struct{}
)

// Mac's collection of methods
func (s *Mac) Print() {
	fmt.Println("Mac print")
	s.printer.PrintFile()
}

func (s *Mac) SetPrinter(printer Printer) {
	s.printer = printer
}

// Windows's collection of methods
func (s *Windows) Print() {
	fmt.Println("Windows print")
	s.printer.PrintFile()
}

func (s *Windows) SetPrinter(printer Printer) {
	s.printer = printer
}

// Epson's collection of methods
func (s *Epson) PrintFile() {
	fmt.Println("Epson print")
}

// HP's collection of methods
func (s *HP) PrintFile() {
	fmt.Println("HP print")
}

// main function
func main() {
	hpPrinter := new(HP)
	epsonPrinter := new(Epson)

	macComputer := new(Mac)
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()

	windowsComputer := new(Windows)
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
}