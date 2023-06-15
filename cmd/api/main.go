package main

import (
	"context"
	"github.com/jumagaliev1/hash-generator/internal/app"
	"github.com/jumagaliev1/hash-generator/internal/config"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := app.New().Run(ctx); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
