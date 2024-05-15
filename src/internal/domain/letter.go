package domain

type Letter struct {
	Value string
}

func NewLetter(v string) *Letter {
	return &Letter{Value: v}
}
