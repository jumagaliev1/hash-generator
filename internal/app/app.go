package app

import (
	"context"
	"github.com/jumagaliev1/hash-generator/internal/service"
	"github.com/jumagaliev1/hash-generator/internal/storage"
	"github.com/jumagaliev1/hash-generator/internal/transport"
)

var addr = ":50051"

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	repo, err := storage.New(ctx)
	if err != nil {
		return err
	}

	svc, err := service.New(repo)
	if err != nil {
		return err
	}

	if err := transport.Run(svc); err != nil {
		return err
	}

	return nil
}
