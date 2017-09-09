package adapter

import (
	"testing"
	"net/http"
)

func TestPrinterAdapter(t *testing.T) {

	msg := "Hello World"
	adapter := PrinterAdapter{&HomeLegacyPrinter{}, msg}

	returnedMsg := adapter.PrintStored()

	if returnedMsg != "LEGACY: Adapter: Hello World" {
		t.Errorf("Message Did not match: %s", returnedMsg)
	}

	adapter2 := PrinterAdapter{nil, msg}

	returnedMsg2 := adapter2.PrintStored()

	if returnedMsg2 != msg {
		t.Errorf("Message Did not match: %s", returnedMsg2)
	}

http.HandleFunc()

}
