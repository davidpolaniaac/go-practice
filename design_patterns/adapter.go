package main

import "fmt"

type Payment interface {
	pay()
}

type PaymentCash struct{}

func (PaymentCash) pay() {
	fmt.Println("Cash payment")
}

func printPayment(p Payment) {
	p.pay()
}

type PaymentBank struct{}

func (PaymentBank) pay(account int) {
	fmt.Printf("Bank payment with account %d \n", account)
}

type PaymentBankAdapter struct {
	account     int
	PaymentBank *PaymentBank
}

func (p *PaymentBankAdapter) pay() {
	p.PaymentBank.pay(p.account)
}

func main() {
	paymentCash := &PaymentCash{}
	printPayment(paymentCash)
	paymentBank := &PaymentBank{}
	//printPayment(paymentBank)

	adapter := &PaymentBankAdapter{
		account:     5,
		PaymentBank: paymentBank,
	}

	printPayment(adapter)

}
