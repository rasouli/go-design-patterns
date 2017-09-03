package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	pm, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("expected a cash payment method but go nil")
	}

	msg := pm.Pay(10.30)

	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("expected pay by cach method, got the message : %s", msg)
	}

	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebit(t *testing.T) {

	pm, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("expected a debit payment method but got nil")
	}

	msg := pm.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Errorf("expected a debit payment message but got %s ", msg)
	}
}

func TestGetPaymentMethodNotExists(t *testing.T) {

	fakeMethod := 20

	_, err := GetPaymentMethod(fakeMethod)

	if err == nil {
		t.Errorf("a fake payment method should not exist among payment method with the id of %d", fakeMethod)
	}
}
