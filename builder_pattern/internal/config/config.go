package config

import "github.com/nurashi/software-design-patterns/builder_pattern/internal/config/subconfigs"

type Config struct {
	OpenRouter subconfigs.OpenRouter
	Database   subconfigs.PostgreSQL
	Telegram   subconfigs.Telegram
}

type ConfigBuilder struct {
	cfg Config
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (b *ConfigBuilder) SetDatabase(db subconfigs.PostgreSQL) *ConfigBuilder {
	b.cfg.Database = db
	return b
}

func (b *ConfigBuilder) SetTelegram(tg subconfigs.Telegram) *ConfigBuilder {
	b.cfg.Telegram = tg
	return b
}

func (b *ConfigBuilder) SetOpenRouter(or subconfigs.OpenRouter) *ConfigBuilder {
	b.cfg.OpenRouter = or
	return b
}

func (b *ConfigBuilder) Build() Config {
	return b.cfg
}