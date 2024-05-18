package domain

type Color string

const Green Color = "green"
const Yellow Color = "yellow"
const Gray Color = "gray"

type Letter struct {
	Value rune
	Color Color
}

func NewLetter(v rune) *Letter {
	return &Letter{Value: v, Color: Gray}
}
