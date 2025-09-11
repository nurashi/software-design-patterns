package subconfigs

// Telegram holds configuration for the Telegram bot
type Telegram struct {
	Token string
}

// TelegramBuilder provides a fluent way to build a Telegram configuration
type TelegramBuilder struct {
	tg Telegram
}

// NewTelegramBuilder returns a new TelegramBuilder
func NewTelegramBuilder() *TelegramBuilder {
	return &TelegramBuilder{}
}

func (b *TelegramBuilder) SetToken(token string) *TelegramBuilder {
	b.tg.Token = token
	return b
}

// Build returns the configured Telegram instance
func (b *TelegramBuilder) Build() Telegram {
	return b.tg
}
