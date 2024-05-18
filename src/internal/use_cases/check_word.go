package usecases

import (
	"context"
	"fmt"
	"wordle/internal/domain"
)

type Repository interface {
	GetDailyWord(ctx context.Context) (*domain.DailyWord, error)
	CheckInDictionaryExists(ctx context.Context, ew *domain.EnteredWord) (bool, error)
}

func CheckWord(ctx context.Context, repo Repository, ew *domain.EnteredWord) (*domain.EnteredWord, error) {
	exists, err := repo.CheckInDictionaryExists(ctx, ew)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("%s", "несуществующее слово")
	}
	dw, err := repo.GetDailyWord(ctx)
	if err != nil {
		return nil, err
	}
	err = ew.CheckLetters(dw)
	if err != nil {
		return nil, err
	}

	return ew, nil
}
