package usecases

import (
	"context"
	"wordle/internal/domain"
	"wordle/internal/domain/dto"
)

type Repository interface {
	GetDailyWord(ctx context.Context) (*domain.Word, error)
}

func CheckWord(ctx context.Context, repo Repository, params dto.CheckWordParams) (*domain.Word, error) {
	w, err := repo.GetDailyWord(ctx)
	if err != nil {
		return nil, err
	}
	err = w.Check(params)
	if err != nil {
		return nil, err
	}

	return w, nil
}
