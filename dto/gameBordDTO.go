package dto

import (
	O "github.com/IBM/fp-go/option"
)

type GameBord struct {
	PartyA Party  `json:"party_a"`
	PartyB Party  `json:"party_b"`
	Winner string `json:"winner"`
}

func (g GameBord) GetPartyA() Party {
	return g.PartyA
}
func (g GameBord) GetPartyB() Party {
	return g.PartyB
}
func (g GameBord) GetWinner() string {
	return g.Winner
}

type Party struct {
	ID       string  `json:"id"`
	Rock     Pokemon `json:"rock"`
	Scissors Pokemon `json:"scissors"`
	Paper    Pokemon `json:"paper"`
	HP       HP      `json:"hp"`
	Jaken    Jaken   `json:"jaken"`
}

type Jaken struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Stats struct {
	HP   int `json:"hp"`
	Atk  int `json:"atk"`
	Df   int `json:"df"`
	SAtk int `json:"satk"`
	SDf  int `json:"sdf"`
}

type Pokemon struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Types   string  `json:"types"`
	Stats   Stats   `json:"stats"`
	Sprites Sprites `json:"sprites"`
}

type HP struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}

type Sprites struct {
	Front string `json:"front"`
}

func NewPokemon(id int, name, types string, stats Stats, sprites Sprites) Pokemon {
	return Pokemon{
		ID:      id,
		Name:    name,
		Types:   types,
		Stats:   stats,
		Sprites: sprites,
	}
}

func NewStats(hp, atk, df, satk, sdf int) Stats {
	return Stats{
		HP:   hp,
		Atk:  atk,
		Df:   df,
		SAtk: satk,
		SDf:  sdf,
	}
}

func NewSprites(front string) Sprites {
	return Sprites{
		Front: front,
	}
}

func NewHP(current, max int) HP {
	return HP{
		Current: current,
		Max:     max,
	}
}

func NewJaken(name string, value int) Jaken {
	return Jaken{
		Name:  name,
		Value: value,
	}
}

func NewParty(id string, rock, scissors, paper Pokemon, hp HP, jaken Jaken) Party {
	return Party{
		ID:       id,
		Rock:     rock,
		Scissors: scissors,
		Paper:    paper,
		HP:       hp,
		Jaken:    jaken,
	}
}

func NewGameBord(partyA, partyB Party, winner O.Option[string]) GameBord {
	return GameBord{
		PartyA: partyA,
		PartyB: partyB,
		Winner: O.Fold(
			func() string { return "" },
			func(name string) string { return name },
		)(winner),
	}
}
