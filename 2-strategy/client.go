package main

import "github.com/nintran52/design-patterns/2-strategy/payment"

func main() {
	context := payment.PaymentContext{}

	context.SetStrategy(payment.CreditCardPayment{})
	context.Checkout(100.0)

	context.SetStrategy(payment.PayPalPayment{})
	context.Checkout(200.5)

	context.SetStrategy(payment.BankTransferPayment{})
	context.Checkout(350.75)
}
