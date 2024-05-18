package main

import (
	"context"
	"fmt"
	"wordle/internal"
	"wordle/internal/domain"
	usecases "wordle/internal/use_cases"
)

func main() {
	ctx := context.Background()
	app, err := internal.NewApp(ctx)
	if err != nil {
		panic(err)
	}
	w := "ПАИТА"
	c, err := usecases.CheckWord(ctx, app.Repo, domain.NewEnteredWord(w))
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%+v\n", c.Letters[i])
		fmt.Println(string(c.Letters[i].Value))
	}
}
