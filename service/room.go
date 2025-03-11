package service

import (
	"my-go-app/api"
	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/model"
	"my-go-app/model/monad/result"
	"my-go-app/util"
)

type Room struct {
	room connection.Room
}

func NewRoomService(room connection.Room) Room {
	return Room{
		room: room,
	}
}

func (s Room) GetPokemon(reqA dto.RegistRequestDto, reqB dto.RegistRequestDto) model.BattleField {
	var res model.BattleField
	res.PlayerA.Name = reqA.Name
	pokemonsA := convertToPokemon(getPokemons(reqA.Paper, reqA.Rock, reqA.Scissors))
	res.PlayerA.Party.Paper = pokemonsA[0]
	res.PlayerA.Party.Rock = pokemonsA[1]
	res.PlayerA.Party.Scissors = pokemonsA[2]
	res.PlayerA.Party.HP.Current = calculateHP(
		res.PlayerA.Party.Paper,
		res.PlayerA.Party.Rock,
		res.PlayerA.Party.Scissors,
	)
	res.PlayerA.Party.HP.Max = res.PlayerA.Party.HP.Current

	res.PlayerB.Name = reqB.Name
	pokemonsB := convertToPokemon(getPokemons(reqB.Paper, reqB.Rock, reqB.Scissors))
	res.PlayerB.Party.Paper = pokemonsB[0]
	res.PlayerB.Party.Rock = pokemonsB[1]
	res.PlayerB.Party.Scissors = pokemonsB[2]
	res.PlayerB.Party.HP.Current = calculateHP(
		res.PlayerB.Party.Paper,
		res.PlayerB.Party.Rock,
		res.PlayerB.Party.Scissors,
	)
	res.PlayerB.Party.HP.Max = res.PlayerB.Party.HP.Current
	return res
}

func (s Room) Jaken(reqA dto.JakenRequestDto, reqB dto.JakenRequestDto, battleField model.BattleField) model.BattleField {
	winner := util.JudgeJaken(reqA, reqB)
	res := util.CalculateDamage(battleField, winner)
	return res
}

func calculateHP(pokemons ...model.Pokemon) int {
	hp := 0
	for _, pokemon := range pokemons {
		hp += pokemon.Stats[0].BaseStat
	}
	return hp
}

func getPokemons(ids ...int) []dto.PokemonResponseDto {
	api := api.PokeAPI{}
	var pokemons []dto.PokemonResponseDto
	for _, id := range ids {
		pokemon := api.GetPokemon(id)
		pokemons = append(pokemons, result.Unwrap[dto.PokemonResponseDto](pokemon))
	}
	return pokemons
}

func convertToPokemon(pokemonDtos []dto.PokemonResponseDto) []model.Pokemon {
	var pokemons []model.Pokemon
	for _, pokemonDto := range pokemonDtos {
		pokemons = append(pokemons, model.Pokemon{
			ID:      pokemonDto.ID,
			Name:    pokemonDto.Name,
			Types:   pokemonDto.Types,
			Stats:   pokemonDto.Stats,
			Sprites: pokemonDto.Sprites,
			Species: pokemonDto.Species,
		})
	}
	return pokemons
}
