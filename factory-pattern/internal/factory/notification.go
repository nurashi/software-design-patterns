package factory

import "github.com/nurashi/software-design-patterns/internal/notification"


// its a contract for factories
// EmailFactory -> EmailNotification
// SMSFactory -> SMSNotification 
// TelegramFactory -> TelegramNotification 
type NotificationFactory interface {
	CreateNotification() notification.Notification
}

type EmailFactory struct{}

func (f EmailFactory) CreateNotification() notification.Notification {
	return notification.EmailNotification{}
}

type SMSFactory struct{}

func (f SMSFactory) CreateNotification() notification.Notification {
	return notification.SMSNotification{}
}

type TelegramFactory struct{}

func (f TelegramFactory) CreateNotification() notification.Notification {
	return notification.TelegramNotification{}
}



