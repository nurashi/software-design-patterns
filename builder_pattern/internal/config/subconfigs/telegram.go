package subconfigs 


type Telegram struct {
	Token string
}


type TelegramBuilder struct {
	tg Telegram
}

func NewTelegramBuilder() *TelegramBuilder {
	return &TelegramBuilder{}
}

func (b *TelegramBuilder) SetToken(token string) *TelegramBuilder {
	b.tg.Token = token
	return b
}

func (b *TelegramBuilder) Build() Telegram {
	return b.tg
}
