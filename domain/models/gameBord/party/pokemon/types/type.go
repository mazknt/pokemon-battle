package types

import (
	"errors"

	E "github.com/IBM/fp-go/either"
)

type pokemonType string

const (
	Normal   pokemonType = "normal"
	Fighting pokemonType = "fighting"
	Flying   pokemonType = "flying"
	Poison   pokemonType = "poison"
	Ground   pokemonType = "ground"
	Rock     pokemonType = "rock"
	Bug      pokemonType = "bug"
	Ghost    pokemonType = "ghost"
	Steel    pokemonType = "steel"
	Fire     pokemonType = "fire"
	Water    pokemonType = "water"
	Grass    pokemonType = "grass"
	Electric pokemonType = "electric"
	Psychic  pokemonType = "psychic"
	Ice      pokemonType = "ice"
	Dragon   pokemonType = "dragon"
	Dark     pokemonType = "dark"
	Fairy    pokemonType = "fairy"
)

var validTypes = map[string]pokemonType{
	"normal":   Normal,
	"fighting": Fighting,
	"flying":   Flying,
	"poison":   Poison,
	"ground":   Ground,
	"rock":     Rock,
	"bug":      Bug,
	"ghost":    Ghost,
	"steel":    Steel,
	"fire":     Fire,
	"water":    Water,
	"grass":    Grass,
	"electric": Electric,
	"psychic":  Psychic,
	"ice":      Ice,
	"dragon":   Dragon,
	"dark":     Dark,
	"fairy":    Fairy,
}

type Types struct {
	value pokemonType
}

// NewTypes PokemonTypes から Types を生成する関数
func New(pt string) E.Either[error, Types] {
	t, ok := validTypes[string(pt)]
	if !ok {
		return E.Left[Types](errors.New("invalid PokemonType"))
	}
	return E.Right[error](Types{value: t})
}

func (t Types) GetValue() pokemonType {
	return t.value
}
