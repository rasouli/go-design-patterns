package bridge

import (
	"testing"
	"strings"
)

func TestPrinterImpl1(t *testing.T) {
	api1 := PrinterImpl1{}

	err := api1.PrintMessage("Hello")

	if err != nil {
		t.Errorf("Error Trying to use the API1 implementation: %s", err.Error())
	}

}

func TestPrinterImpl2(t *testing.T) {

	api2 := PrinterImpl2{}

	err := api2.PrintMessage("Hello")

	if err != nil {
		expectedErrorMsg := "you need to pass an io.Writer to PrinterImpl2"

		if !strings.Contains(err.Error(), expectedErrorMsg) {
			t.Errorf("Expected Error Message Was Not %s , got %s", expectedErrorMsg, err.Error())
		}
	}

	testWriter := TestWriter{}
	api2 = PrinterImpl2{&testWriter}
	expectedMsg := "Hello"

	err = api2.PrintMessage("Hello")

	if err != nil {
		t.Errorf("failed to use api v2: %s ", err.Error())
	}

	if testWriter.Msg != expectedMsg {
		t.Errorf("Expected message to be %s but got %s", expectedMsg, testWriter.Msg)
	}

}

func TestNormalPrinter_Print(t *testing.T) {
	expectedMessage := "Hello io.Writer"

	normalPrinter := NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl1{},
	}

	err := normalPrinter.Print()
	if err != nil {
		t.Errorf("failed to print with the normal printer: %s", err.Error())
	}

	testWriter := TestWriter{}
	normalPrinter = NormalPrinter{
		Msg:     expectedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = normalPrinter.Print()
	if err != nil {
		t.Errorf("failed to print with the normal printer: %s", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("the message on the test writer does not match the actual. expected: %s, Actual: %s", expectedMessage, testWriter.Msg)
	}
}

func TestPacktPrinter_Print(t *testing.T) {
	passedMessage := "Hello io.Writer"
	expectedMessage := "Message from Packt: Hello io.Writer"


	packt := PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl1{},
	}

	err := packt.Print()
	if err != nil {
		t.Errorf("failed to print with the packt printer: %s", err.Error())
	}

	testWriter := TestWriter{}
	packt = PacktPrinter{
		Msg:     passedMessage,
		Printer: &PrinterImpl2{Writer: &testWriter},
	}

	err = packt.Print()
	if err != nil {
		t.Errorf("failed to print with the normal printer: %s", err.Error())
	}

	if testWriter.Msg != expectedMessage {
		t.Errorf("the message on the test writer does not match the actual. expected: %s, Actual: %s", expectedMessage, testWriter.Msg)
	}
}