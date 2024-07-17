package main

import (
	"bank/internal/app"
	"bank/internal/config"
)

func main() {
	cfg := config.MustNew()
	app.Run(cfg)
}
