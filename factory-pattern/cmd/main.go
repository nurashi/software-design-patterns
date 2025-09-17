package main

import (
	"fmt"

	"github.com/nurashi/software-design-patterns/internal/factory"
)

func main() {

	// client code
	var f factory.NotificationFactory

	f = factory.EmailFactory{}
	email := f.CreateNotification()
	fmt.Println(email.Send("nurashi@gmail.com", "Salem"))

	f = factory.SMSFactory{}
	sms := f.CreateNotification()
	fmt.Println(sms.Send("+77777777777", "Your code is 7777"))

	f = factory.TelegramFactory{}
	tg := f.CreateNotification()
	fmt.Println(tg.Send("@nurashi", "Hello"))
}


/*
           ┌─────────────────────────┐
           │    Notification         │
           │  + Send(to,msg):string  │
           └───────────▲────────────┘
                       │
 ┌─────────────────────┼──────────────────────┐
 │                     │                      │
 │                     │                      │
 ▼                     ▼                      ▼
EmailNotification   SMSNotification    TelegramNotification


         ┌────────────────────────────┐
         │   NotificationFactory      │
         │  + CreateNotification()    │
         └───────────▲────────────────┘
                     │
   ┌─────────────────┼───────────────────┐
   │                 │                   │
   ▼                 ▼                   ▼
EmailFactory      SMSFactory       TelegramFactory
*/