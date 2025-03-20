package jaken

import (
	"errors"

	E "github.com/IBM/fp-go/either"
)

type Jaken struct {
	name  string
	value int
}

func (j Jaken) Name() string {
	return j.name
}
func (j Jaken) Value() int {
	return j.value
}

func New(name string) E.Either[error, Jaken] {
	if name == "rock" {
		return E.Right[error](NewRock())
	}
	if name == "scissors" {
		return E.Right[error](NewScissors())
	}
	if name == "paper" {
		return E.Right[error](NewPaper())
	}
	return E.Left[Jaken](errors.New("only rock or paper or scissors"))
}

func NewRock() Jaken {
	return Jaken{
		name:  "rock",
		value: 0,
	}
}
func NewScissors() Jaken {
	return Jaken{
		name:  "scissors",
		value: 1,
	}
}
func NewPaper() Jaken {
	return Jaken{
		name:  "paper",
		value: 2,
	}
}
