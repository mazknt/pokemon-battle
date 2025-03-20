package service

import (
	jakenDomainService "my-go-app/domain/domainService/jaken"
	partyDomainService "my-go-app/domain/domainService/party"
	gamebord "my-go-app/domain/models/gameBord"
	"my-go-app/domain/models/gameBord/jaken"
	"my-go-app/domain/models/gameBord/party"
	"my-go-app/domain/models/gameBord/party/pokemon"
	"my-go-app/dto"
	"my-go-app/dto/mapper"
	apiif "my-go-app/service/apiIF"

	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
)

func NewGameBord(pokeAPI apiif.PokeAPI) GameBordService {
	return GameBordService{
		pokeAPI: pokeAPI,
	}
}

type GameBordService struct {
	pokeAPI apiif.PokeAPI
}

func (g GameBordService) CreateGameBord(reqA dto.RegistRequestDto) func(reqB dto.RegistRequestDto) E.Either[error, dto.GameBord] {
	return func(reqB dto.RegistRequestDto) E.Either[error, dto.GameBord] {
		return F.Pipe4(
			E.Right[error](gamebord.NewGameBord),
			E.Ap[func(partyB party.Party) func(winner O.Option[string]) gamebord.GameBord](g.partyFromRegistDTO(reqA)),
			E.Ap[func(winner O.Option[string]) gamebord.GameBord](g.partyFromRegistDTO(reqB)),
			E.Ap[gamebord.GameBord](E.Right[error](O.None[string]())),
			E.Map[error](mapper.GameBordToDTO),
		)
	}
}

func (g GameBordService) Battle(battleField dto.GameBord) func(playerA dto.JakenRequestDto) func(playerB dto.JakenRequestDto) E.Either[error, dto.GameBord] {
	return func(playerA dto.JakenRequestDto) func(playerB dto.JakenRequestDto) E.Either[error, dto.GameBord] {
		return func(playerB dto.JakenRequestDto) E.Either[error, dto.GameBord] {
			gamebdE := mapper.GameBordFromDTO(battleField)
			handE := winnerJaken(playerA)(playerB)
			winnerE := F.Pipe2(
				E.Right[error](winnerParty(playerA)(playerB)),
				E.Ap[func(hand jaken.Jaken) party.Party](gamebdE),
				E.Ap[party.Party](handE),
			)
			loserE := E.Chain[error](
				func(winner party.Party) E.Either[error, party.Party] {
					return E.Map[error](
						func(gamebd gamebord.GameBord) party.Party {
							return gamebd.GetOtherParty(winner)
						},
					)(gamebdE)
				})(winnerE)
			return F.Pipe7(
				E.Right[error](partyDomainService.CalculateDamage),
				E.Ap[func(loser party.Party) func(hand jaken.Jaken) E.Either[error, party.Party]](winnerE),
				E.Ap[func(hand jaken.Jaken) E.Either[error, party.Party]](loserE),
				E.Ap[E.Either[error, party.Party]](handE),
				E.Flatten,
				E.Chain(
					func(damagedLooser party.Party) E.Either[error, gamebord.GameBord] {
						return E.Map[error](
							func(gamebd gamebord.GameBord) gamebord.GameBord {
								return gamebd.UpdateParty(damagedLooser)
							},
						)(gamebdE)
					}),
				judgeLooser(playerA.Name)(playerB.Name),
				E.Map[error](mapper.GameBordToDTO),
			)
		}
	}
}

func judgeLooser(playerA string) func(playerB string) func(gamebdE E.Either[error, gamebord.GameBord]) E.Either[error, gamebord.GameBord] {
	return func(playerB string) func(gamebdE E.Either[error, gamebord.GameBord]) E.Either[error, gamebord.GameBord] {
		return func(gamebdE E.Either[error, gamebord.GameBord]) E.Either[error, gamebord.GameBord] {
			return E.Fold(
				func(err error) E.Either[error, gamebord.GameBord] {
					if err.Error() == playerA || err.Error() == playerB {
						return E.Right[error](gamebord.ResultBattle(err.Error()))
					}
					return E.Left[gamebord.GameBord](err)
				},
				func(gamebd gamebord.GameBord) E.Either[error, gamebord.GameBord] {
					return E.Right[error](gamebd)
				},
			)(gamebdE)
		}
	}

}

func (g GameBordService) partyFromRegistDTO(regist dto.RegistRequestDto) E.Either[error, party.Party] {
	return F.Pipe3(
		E.Right[error](party.Init(regist.Name)),
		E.Ap[func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) party.Party](E.Chain[error](mapper.PokemonFromDTO)(g.pokeAPI.GetPokemon(regist.Rock))),
		E.Ap[func(scissors pokemon.Pokemon) party.Party](E.Chain[error](mapper.PokemonFromDTO)(g.pokeAPI.GetPokemon(regist.Paper))),
		E.Ap[party.Party](E.Chain[error](mapper.PokemonFromDTO)(g.pokeAPI.GetPokemon(regist.Scissors))),
	)
}

func winnerJaken(playerA dto.JakenRequestDto) func(playerB dto.JakenRequestDto) E.Either[error, jaken.Jaken] {
	return func(playerB dto.JakenRequestDto) E.Either[error, jaken.Jaken] {
		return F.Pipe3(
			E.Right[error](jakenDomainService.DetermineJakenWinnerDomainService),
			E.Ap[func(jakenB jaken.Jaken) E.Either[error, jaken.Jaken]](convertToEntityFromDTO(playerA)),
			E.Ap[E.Either[error, jaken.Jaken]](convertToEntityFromDTO(playerB)),
			E.Flatten,
		)
	}
}

func winnerParty(playerA dto.JakenRequestDto) func(playerB dto.JakenRequestDto) func(gamebd gamebord.GameBord) func(hand jaken.Jaken) party.Party {
	return func(playerB dto.JakenRequestDto) func(gamebd gamebord.GameBord) func(hand jaken.Jaken) party.Party {
		return func(gamebd gamebord.GameBord) func(hand jaken.Jaken) party.Party {
			return func(hand jaken.Jaken) party.Party {
				return F.Pipe2(
					hand,
					getWinnerID(playerA)(playerB),
					gamebd.GetParty,
				)
			}
		}
	}
}

func getWinnerID(reqA dto.JakenRequestDto) func(reqB dto.JakenRequestDto) func(hand jaken.Jaken) string {
	return func(reqB dto.JakenRequestDto) func(hand jaken.Jaken) string {
		return func(hand jaken.Jaken) string {
			if hand.Name() == reqA.Jaken.Name {
				return reqA.Name
			}
			return reqB.Name
		}
	}
}

func convertToEntityFromDTO(jakenRequest dto.JakenRequestDto) E.Either[error, jaken.Jaken] {
	return jaken.New(jakenRequest.Jaken.Name)
}
