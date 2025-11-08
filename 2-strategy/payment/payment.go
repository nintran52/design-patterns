package payment

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct{}

func (CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Credit Card\n", amount)
}

type PayPalPayment struct{}

func (PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paying %.2f using PayPal\n", amount)
}

type BankTransferPayment struct{}

func (BankTransferPayment) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Bank Transfer\n", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (c *PaymentContext) SetStrategy(s PaymentStrategy) {
	c.strategy = s
}

func (c *PaymentContext) Checkout(amount float64) {
	c.strategy.Pay(amount)
}
