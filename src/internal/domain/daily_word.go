package domain

import (
	"unicode/utf8"
)

type DailyWord struct {
	Letters [5]*Letter
	Count   map[rune]int
}

func NewDailyWord(dw string) *DailyWord {
	var letters [5]*Letter
	letterCount := make(map[rune]int, len(dw))
	runes := []rune(dw)
	for i := 0; i < utf8.RuneCountInString(dw); i++ {
		letters[i] = NewLetter(runes[i])
		letterCount[runes[i]]++
	}

	return &DailyWord{Letters: letters, Count: letterCount}
}
