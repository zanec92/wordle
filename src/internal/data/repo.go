package data

import (
	"context"
	"wordle/internal/domain"
)

const dailyWord = "пицца"

type Repository struct{}

func New(ctx context.Context) (*Repository, error) {
	return &Repository{}, nil
}

func (r Repository) GetDailyWord(ctx context.Context) (*domain.DailyWord, error) {
	return domain.NewDailyWord(dailyWord), nil
}

func (r Repository) CheckInDictionaryExists(ctx context.Context, ew *domain.EnteredWord) (bool, error) {
	return true, nil
}
