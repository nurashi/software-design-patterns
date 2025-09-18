
// this package will represent main structs, based on this structs will be created objects. 
package notification

import "fmt"


// main Email struct
type EmailNotification struct{}

func (e EmailNotification) Send(to string, message string) string {
	return fmt.Sprintf("Email to %s: %s", to, message)
}

// main sms struct 
type SMSNotification struct{}

func (s SMSNotification) Send(to string, message string) string {
	return fmt.Sprintf("SMS to %s: %s", to, message)
}

// main telegram struct 
type TelegramNotification struct{}

func (t TelegramNotification) Send(to string, message string) string {
	return fmt.Sprintf("Telegram to %s: %s", to, message)
}
