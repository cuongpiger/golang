package adapter

import "fmt"

type (
	LegacyPrinter interface {
		Print(s string) string
	}

	NewPrinter interface {
		PrintStored() string
	}
)

type (
	MyLegacyPrinter struct{}
	PrinterAdapter  struct {
		OldPrinter LegacyPrinter
		Msg        string
	}
)

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	println(newMsg)
	return
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}

	return
}
