package pokemondata

import (
	"my-go-app/domain/models/gameBord/party/pokemon"
	"my-go-app/domain/models/gameBord/party/pokemon/sprites"
	"my-go-app/domain/models/gameBord/party/pokemon/stats"
	"my-go-app/domain/models/gameBord/party/pokemon/types"

	E "github.com/IBM/fp-go/either"
)

var TestData1 pokemon.Pokemon = makePokemonData(1, "Bulbasaur", "grass", stat{hp: 45, atk: 49, df: 49, satk: 65, sdf: 65})
var TestData2 pokemon.Pokemon = makePokemonData(4, "Charmander", "fire", stat{hp: 39, atk: 52, df: 43, satk: 60, sdf: 50})
var TestData3 pokemon.Pokemon = makePokemonData(7, "Squirtle", "water", stat{hp: 44, atk: 48, df: 65, satk: 50, sdf: 64})
var TestData4 pokemon.Pokemon = makePokemonData(25, "Pikachu", "electric", stat{hp: 35, atk: 55, df: 40, satk: 50, sdf: 50})
var TestData5 pokemon.Pokemon = makePokemonData(39, "Jigglypuff", "fairy", stat{hp: 115, atk: 45, df: 20, satk: 45, sdf: 25})
var TestData6 pokemon.Pokemon = makePokemonData(150, "Mewtwo", "psychic", stat{hp: 106, atk: 110, df: 90, satk: 154, sdf: 90})

func makePokemonData(id int, name string, typ string, sts stat) pokemon.Pokemon {
	typE := types.New(typ)
	ty, _ := E.Unwrap(typE)
	spr := sprites.New("test")
	status := stats.NewStats(sts.hp, sts.atk, sts.df, sts.satk, sts.sdf)
	return pokemon.New(id)(name)(ty)(status)(spr)
}

type stat struct {
	hp   int
	atk  int
	df   int
	satk int
	sdf  int
}
