package partyDomainService

import (
	pokemonDomainService "my-go-app/domain/domainService/party/pokemon"
	"my-go-app/domain/models/gameBord/jaken"
	"my-go-app/domain/models/gameBord/party"
	"my-go-app/domain/models/gameBord/party/pokemon"

	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
)

func getPokemonFromJaken(p party.Party) func(hand jaken.Jaken) pokemon.Pokemon {
	return func(hand jaken.Jaken) pokemon.Pokemon {
		var pokemon pokemon.Pokemon
		if hand.Name() == "rock" {
			pokemon = p.Rock()
		}
		if hand.Name() == "scissors" {
			pokemon = p.Scissors()
		}
		if hand.Name() == "paper" {
			pokemon = p.Paper()
		}
		return pokemon
	}
}

func CalculateDamage(winner party.Party) func(loser party.Party) func(hand jaken.Jaken) E.Either[error, party.Party] {
	return func(loser party.Party) func(hand jaken.Jaken) E.Either[error, party.Party] {
		return func(hand jaken.Jaken) E.Either[error, party.Party] {
			return F.Pipe3(
				E.Right[error](pokemonDomainService.CalculateDamage),
				E.Ap[func(loser pokemon.Pokemon) int](E.Right[error](getPokemonFromJaken(winner)(hand))),
				E.Ap[int](E.Right[error](getPokemonFromJaken(loser)(hand))),
				E.Chain(loser.Damaged),
			)
		}
	}
}
