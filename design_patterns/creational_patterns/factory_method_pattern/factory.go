package factorymethodpattern

import (
	"errors"
	"fmt"
)

// implement a payment method factory which is going to procvide us with different ways of paying at a shop
// In the beginning we will have two methods of paying cash and credit card. We'll also have an interface witht the method pay
// which every struct that wants to be used as a payment method must implement

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(NewDebitCardPM), nil

	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}

type DebitCardPM struct{}
type NewDebitCardPM struct{}

func (d *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card (new)\n", amount)
}
func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f paid using debit card\n", amount)
}
