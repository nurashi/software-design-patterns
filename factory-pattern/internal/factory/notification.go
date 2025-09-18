package factory

import "github.com/nurashi/software-design-patterns/internal/notification"


// its a contract for factories
type NotificationFactory interface {
	CreateNotification() notification.Notification
}

// EmailFactory -> EmailNotification
type EmailFactory struct{}

func (f EmailFactory) CreateNotification() notification.Notification {
	return notification.EmailNotification{}
}

// SMSFactory -> SMSNotification 
type SMSFactory struct{}

func (f SMSFactory) CreateNotification() notification.Notification {
	return notification.SMSNotification{}
}

// TelegramFactory -> TelegramNotification 
type TelegramFactory struct{}

func (f TelegramFactory) CreateNotification() notification.Notification {
	return notification.TelegramNotification{}
}



