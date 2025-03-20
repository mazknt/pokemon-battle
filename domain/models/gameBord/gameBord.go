package gamebord

import (
	"errors"
	"my-go-app/domain/models/gameBord/jaken"
	"my-go-app/domain/models/gameBord/party"
	"my-go-app/domain/models/gameBord/party/hp"
	"my-go-app/domain/models/gameBord/party/pokemon"

	E "github.com/IBM/fp-go/either"
	O "github.com/IBM/fp-go/option"
)

type GameBord struct {
	partyA party.Party
	partyB party.Party
	winner O.Option[string]
}

func NewGameBord(partyA party.Party) func(partyB party.Party) func(winner O.Option[string]) GameBord {
	return func(partyB party.Party) func(winner O.Option[string]) GameBord {
		return func(winner O.Option[string]) GameBord {
			return GameBord{
				partyA: partyA,
				partyB: partyB,
				winner: winner,
			}
		}
	}
}

func ResultBattle(looser string) GameBord {
	return GameBord{
		partyA: party.Party{},
		partyB: party.Party{},
		winner: O.Some(looser),
	}
}

func (g GameBord) UpdateWinner(winner string) E.Either[error, GameBord] {
	if winner == "" || winner != g.partyA.ID() || winner != g.partyB.ID() {
		return E.Left[GameBord](errors.New("勝者の名前が不適切です。"))
	}
	return E.Right[error](NewGameBord(g.partyA)(g.partyB)(O.Some(winner)))
}

func (g GameBord) UpdateParty(damagedLooser party.Party) GameBord {
	if g.partyA.ID() == damagedLooser.ID() {
		return NewGameBord(damagedLooser)(g.partyB)(O.None[string]())
	}
	return NewGameBord(g.partyA)(damagedLooser)(O.None[string]())
}

func (g GameBord) PartyA() party.Party {
	return g.partyA
}
func (g GameBord) PartyB() party.Party {
	return g.partyB
}
func (g GameBord) Winner() O.Option[string] {
	return g.winner
}
func (g GameBord) IDOfA() string {
	return g.partyA.ID()
}
func (g GameBord) IDOfB() string {
	return g.partyB.ID()
}
func (g GameBord) RockOfA() pokemon.Pokemon {
	return g.partyA.Rock()
}
func (g GameBord) RockOfB() pokemon.Pokemon {
	return g.partyB.Rock()
}
func (g GameBord) ScissorsOfA() pokemon.Pokemon {
	return g.partyA.Scissors()
}
func (g GameBord) ScissorsOfB() pokemon.Pokemon {
	return g.partyB.Scissors()
}
func (g GameBord) PaperOfA() pokemon.Pokemon {
	return g.partyA.Paper()
}
func (g GameBord) PaperOfB() pokemon.Pokemon {
	return g.partyB.Paper()
}
func (g GameBord) HPOfA() hp.HP {
	return g.partyA.HP()
}
func (g GameBord) HPOfB() hp.HP {
	return g.partyB.HP()
}
func (g GameBord) JakenOfA() jaken.Jaken {
	return g.partyA.Jaken()
}
func (g GameBord) JakenOfB() jaken.Jaken {
	return g.partyB.Jaken()
}

func (g GameBord) GetParty(name string) party.Party {
	if g.partyA.ID() == name {
		return g.partyA
	}
	return g.partyB
}
func (g GameBord) GetOtherParty(pty party.Party) party.Party {
	if g.partyA == pty {
		return g.partyB
	}
	return g.partyA
}
