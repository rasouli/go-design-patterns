package bridge

import (
	"errors"
	"io"
	"fmt"
)

////////////////////////////////////////
type PrinterApi interface {
	PrintMessage(string) error
}
////////////////////////////////////////
type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s \n", msg)
	return nil
}
////////////////////////////////////////
type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}else {
		fmt.Fprintf(p.Writer,"%s", msg)
		return nil
	}
}

////////////////////////////////////////
type TestWriter struct{
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}

	return n, errors.New("received empty content")
}


/////////////////////////////////////////
type NormalPrinter struct {
	Msg string
	Printer PrinterApi
}

func (np *NormalPrinter) Print() error {
	return np.Printer.PrintMessage(np.Msg)
}

//////////////////////////////////////////
type PacktPrinter struct{
	Msg string
	Printer PrinterApi
}

func (c *PacktPrinter) Print() error {
	return c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
}