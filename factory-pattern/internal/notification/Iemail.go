package notification



// its a contract, that will manage all notifications, Email, Telegram, SMS will use this.
type Notification interface {
	Send(to string, message string) string
}
