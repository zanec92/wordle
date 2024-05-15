package main

import (
	"context"
	"fmt"
	"wordle/internal"
	"wordle/internal/domain/dto"
	usecases "wordle/internal/use_cases"
)

func main() {
	ctx := context.Background()
	app, err := internal.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(usecases.CheckWord(ctx, app.Repo, dto.CheckWordParams{}))
}
