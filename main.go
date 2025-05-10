package main

import "fmt"

// interface
type Notifier interface {
	Send(message string)
}

type EmailNotification struct {
	EmailAddress string
}

func (e EmailNotification) Send(message string) {
	fmt.Printf("📧 Emailga [%s] yuborildi: %s\n", e.EmailAddress, message)
}

type SMSNotification struct {
	PhoneNumber string
}

func (s SMSNotification) Send(message string) {
	fmt.Printf("📱 SMS raqam [%s] ga yuborildi: %s\n", s.PhoneNumber, message)
}

func sendNotification(n Notifier, msg string) {
	n.Send(msg)
}

func main() {
	email := EmailNotification{EmailAddress: "user@example.com"}
	sms := SMSNotification{PhoneNumber: "+998901234567"}

	sendNotification(email, "Salom, bu email xabarnoma!")
	sendNotification(sms, "Salom, bu SMS xabarnoma!")
}
