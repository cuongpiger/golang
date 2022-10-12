package bridge

import (
	"errors"
	"fmt"
	"io"
)

type (
	PrinterAPI interface {
		PrintMessage(string) error
	}
	PrinterAbstraction interface {
		Print() error
	}
)

type (
	PrinterAPI1 struct{}
	PrinterAPI2 struct {
		Writer io.Writer
	}
	NormalPrinter struct {
		Msg     string
		Printer PrinterAPI
	}
	PacktPrinter struct {
		Msg     string
		Printer PrinterAPI
	}
	TestWriter struct {
		Msg string
	}
)

func (d *PrinterAPI1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

func (d *PrinterAPI2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterAPI2")
	}

	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}

	err = errors.New("Content received on Writer was empty")
	return
}
