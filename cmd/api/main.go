package main

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/hash-generator/internal/app"
	"github.com/jumagaliev1/hash-generator/internal/config"
	"log"
	"runtime"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	runtime.GOMAXPROCS(6)
	fmt.Println(runtime.NumCPU())
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	application := app.New()
	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
