package adapter

import "fmt"

type LegacyPrinter interface {
	Print(s string) string
}

type HomeLegacyPrinter struct{}

func (hlp *HomeLegacyPrinter) Print(s string) string {
	newMsg := fmt.Sprintf("LEGACY: %s", s)
	println(newMsg)
	return newMsg
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct{
	OldPrinter LegacyPrinter
	Msg string
}

func (p *PrinterAdapter) PrintStored() string {

	if p.OldPrinter != nil {
		newMsg := fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
		return newMsg
	} else {
		return p.Msg
	}
}