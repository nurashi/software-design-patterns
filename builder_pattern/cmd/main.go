package main

import (
    "fmt"
    "github.com/nurashi/software-design-patterns/builder_pattern/internal/config"
)

func main() {
	
    builder := config.NewConfigBuilder()
    director := &config.Director{}
    director.SetConfigBuilder(builder)
    cfg := director.Construct()
    fmt.Println(cfg)
}