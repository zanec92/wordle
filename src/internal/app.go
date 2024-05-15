package internal

import (
	"context"
	"wordle/internal/data"
)

type App struct {
	Repo *data.Repository
}

func NewApp(ctx context.Context) (*App, error) {
	repo, err := data.New(ctx)
	if err != nil {
		return nil, err
	}
	return &App{Repo: repo}, nil
}
