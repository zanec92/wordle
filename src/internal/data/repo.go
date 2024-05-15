package data

import (
	"context"
	"wordle/internal/domain"
)

const dailyWord = "ПИЦЦА"

type Repository struct{}

func New(ctx context.Context) (*Repository, error) {
	return &Repository{}, nil
}

func (r Repository) GetDailyWord(ctx context.Context) (*domain.Word, error) {
	return domain.NewWord(dailyWord), nil
}
