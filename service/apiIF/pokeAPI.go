package apiif

import (
	"my-go-app/dto"

	E "github.com/IBM/fp-go/either"
)

type PokeAPI interface {
	GetPokemon(id int) E.Either[error, dto.PokemonResponseDto]
}
