package domain

import (
	"wordle/internal/domain"
	"wordle/internal/domain/dto"
)

type Word struct {
	Letters []*Letter
	Count   map[string]int
}

func NewWord(w string) *Word {
	letters := make([]*Letter, len(w))
	letterCount := make(map[string]int, len(w))
	for i := 0; i < len(w); i++ {
		letters = append(letters, domain.NewLetter(w[i]))
		letterCount[domain.NewLetter(w[i])]++
	}
	return &Word{Letters: letters}
}

func (w *Word) Check(params dto.CheckWordParams) error {
	for i := 0; i < len(params.Word); i++ {
		l := string(params.Word[i])
		if l == w.Letters[i].Value {
			// подсветка зеленая
		}
		for j := i; j < len(params.Word); j++ {
			if l == w.Letters[j].Value {
				// подсветка желтая
			}
		}
		// подсветка серая
	}

	//собираем буквы map[string]int
	// letters["и"]:2
	// letters["и"]:1
	// if (letter[value] == l) - подсветка желтая

	return nil
}
