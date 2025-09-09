package main

import (
	"fmt"

	"github.com/nurashi/software-design-patterns/builder_pattern/internal/config"
	"github.com/nurashi/software-design-patterns/builder_pattern/internal/config/subconfigs"
)

func main() {
	db := subconfigs.NewPostgreSQLBuilder().
		SetHost("localhost").
		SetPort(5432).
		SetUser("posgres").
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

	cfg := config.NewConfigBuilder().
		SetDatabase(db).
		SetTelegram(tg).
		SetOpenRouter(or).
		Build()
	
	fmt.Println(cfg)
}
