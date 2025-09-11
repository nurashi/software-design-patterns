package config

import "github.com/nurashi/software-design-patterns/builder_pattern/internal/config/subconfigs"

// Director defines the construction process for a Config using a provided ConfigBuilder
type Director struct {
	configBuilder *ConfigBuilder
}

// SetConfigBuilder assigns a ConfigBuilder to the Director
func (d *Director) SetConfigBuilder(builder *ConfigBuilder) *Director {
	d.configBuilder = builder
	return d
}

// Construct builds a Config instance using predefined defaults for PostgreSQL, Telegram and OpenRouter
func (d *Director) Construct() *Config {
	db := subconfigs.NewPostgreSQLBuilder().
		SetHost("localhost").
		SetPort(5432).
		SetUser("postgres").
		SetPassword("password").
		SetName("postgres").
		Build()

	tg := subconfigs.NewTelegramBuilder().
		SetToken("bot token").
		Build()

	or := subconfigs.NewOpenRouterBuilder().
		SetModel("mistralai/mistral-7b-instruct:free").
		SetReferer("https://github.com/nurashi/Newton").
		Build()

	config := d.configBuilder.
		SetDatabase(db).
		SetTelegram(tg).
		SetOpenRouter(or).
		Build()

	return &config
}
