package notification

import "fmt"

// Concrete product
type EmailNotification struct{}

func (e EmailNotification) Send(to string, message string) string {
	return fmt.Sprintf("Email to %s: %s", to, message)
}

type SMSNotification struct{}

func (s SMSNotification) Send(to string, message string) string {
	return fmt.Sprintf("SMS to %s: %s", to, message)
}

type TelegramNotification struct{}

func (t TelegramNotification) Send(to string, message string) string {
	return fmt.Sprintf("Telegram to %s: %s", to, message)
}
