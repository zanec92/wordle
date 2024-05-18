package domain

import (
	"strings"
	"unicode/utf8"
)

type EnteredWord struct {
	Letters [5]*Letter
}

func NewEnteredWord(ew string) *EnteredWord {
	ew = strings.ToLower(ew)
	var letters [5]*Letter
	runes := []rune(ew)
	for i := 0; i < utf8.RuneCountInString(ew); i++ {
		letters[i] = NewLetter(runes[i])
	}

	return &EnteredWord{Letters: letters}
}

func (ew *EnteredWord) CheckLetters(dw *DailyWord) error {
	for i := 0; i < 5; i++ {
		ew.Letters[i].Color = Gray
		if ew.Letters[i].Value == dw.Letters[i].Value {
			ew.Letters[i].Color = Green
			dw.Count[ew.Letters[i].Value]--
		}
	}
	for i := 0; i < 5; i++ {
		if dw.Count[ew.Letters[i].Value] > 0 {
			for j := 0; j < 5; j++ {
				if ew.Letters[i].Value == dw.Letters[j].Value {
					ew.Letters[i].Color = Yellow
					dw.Count[ew.Letters[i].Value]--
					continue
				}
			}
		}
	}

	return nil
}
