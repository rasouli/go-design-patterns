package factory

import (
	"errors"
	"fmt"
)

const (
	Cash = iota
	DebitCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPayment), nil

	case DebitCard:
		return new(DebitPayment), nil

	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recogonised", m))
	}
}

type CashPayment struct{}
type DebitPayment struct{}

func (c *CashPayment) Pay(amount float32) string  { return fmt.Sprintf("%0.2f $ paid using cash") }
func (d *DebitPayment) Pay(amount float32) string { return fmt.Sprintf("%0.2f $ paid using debit card") }
