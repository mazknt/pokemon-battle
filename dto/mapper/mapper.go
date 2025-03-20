package mapper

import (
	"errors"
	gamebord "my-go-app/domain/models/gameBord"
	"my-go-app/domain/models/gameBord/party"
	"my-go-app/domain/models/gameBord/party/hp"
	"my-go-app/domain/models/gameBord/party/pokemon"
	"my-go-app/domain/models/gameBord/party/pokemon/sprites"
	"my-go-app/domain/models/gameBord/party/pokemon/stats"
	"my-go-app/domain/models/gameBord/party/pokemon/types"
	"my-go-app/dto"

	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
)

func PokemonFromDTO(dto dto.PokemonResponseDto) E.Either[error, pokemon.Pokemon] {
	if len(dto.Types) == 0 {
		return E.Left[pokemon.Pokemon](errors.New("types cannot be empty"))
	}
	if len(dto.Stats) < 5 {
		return E.Left[pokemon.Pokemon](errors.New("invalid stats data"))
	}

	typeE := types.New(string(dto.Types[0].Type.Name))

	sts := stats.NewStats(
		dto.Stats[0].BaseStat, // HP
		dto.Stats[1].BaseStat, // Atk
		dto.Stats[2].BaseStat, // Df
		dto.Stats[3].BaseStat, // Satk
		dto.Stats[4].BaseStat, // Sdf
	)

	return E.Chain(
		func(typ types.Types) E.Either[error, pokemon.Pokemon] {
			return E.Right[error](pokemon.New(dto.ID)(dto.Name)(typ)(sts)(sprites.New(dto.Sprites.FrontDefault)))
		})(typeE)
}

func GameBordToDTO(gb gamebord.GameBord) dto.GameBord {
	return dto.NewGameBord(
		dto.NewParty(gb.IDOfA(),
			dto.NewPokemon(
				gb.RockOfA().ID(),
				gb.RockOfA().Name(),
				string(gb.RockOfA().Types().GetValue()),
				dto.NewStats(gb.RockOfA().Stats().GetHP(),
					gb.RockOfA().Stats().GetAtk(),
					gb.RockOfA().Stats().GetDF(),
					gb.RockOfA().Stats().GetSatk(),
					gb.RockOfA().Stats().GetSdf()),
				dto.NewSprites(gb.RockOfA().Sprites().Front())),
			dto.NewPokemon(
				gb.ScissorsOfA().ID(),
				gb.ScissorsOfA().Name(),
				string(gb.ScissorsOfA().Types().GetValue()),
				dto.NewStats(gb.ScissorsOfA().Stats().GetHP(),
					gb.ScissorsOfA().Stats().GetAtk(),
					gb.ScissorsOfA().Stats().GetDF(),
					gb.ScissorsOfA().Stats().GetSatk(),
					gb.ScissorsOfA().Stats().GetSdf()),
				dto.NewSprites(gb.ScissorsOfA().Sprites().Front())),
			dto.NewPokemon(
				gb.PaperOfA().ID(),
				gb.PaperOfA().Name(),
				string(gb.PaperOfA().Types().GetValue()),
				dto.NewStats(gb.PaperOfA().Stats().GetHP(),
					gb.PaperOfA().Stats().GetAtk(),
					gb.PaperOfA().Stats().GetDF(),
					gb.PaperOfA().Stats().GetSatk(),
					gb.PaperOfA().Stats().GetSdf()),
				dto.NewSprites(gb.PaperOfA().Sprites().Front())),
			dto.NewHP(gb.HPOfA().Current(), gb.HPOfA().Max()),
			dto.NewJaken(gb.JakenOfA().Name(), gb.JakenOfA().Value())),
		dto.NewParty(gb.IDOfB(),
			dto.NewPokemon(
				gb.RockOfB().ID(),
				gb.RockOfB().Name(),
				string(gb.RockOfB().Types().GetValue()),
				dto.NewStats(gb.RockOfB().Stats().GetHP(),
					gb.RockOfB().Stats().GetAtk(),
					gb.RockOfB().Stats().GetDF(),
					gb.RockOfB().Stats().GetSatk(),
					gb.RockOfB().Stats().GetSdf()),
				dto.NewSprites(gb.RockOfB().Sprites().Front())),
			dto.NewPokemon(
				gb.ScissorsOfB().ID(),
				gb.ScissorsOfB().Name(),
				string(gb.ScissorsOfB().Types().GetValue()),
				dto.NewStats(gb.ScissorsOfB().Stats().GetHP(),
					gb.ScissorsOfB().Stats().GetAtk(),
					gb.ScissorsOfB().Stats().GetDF(),
					gb.ScissorsOfB().Stats().GetSatk(),
					gb.ScissorsOfB().Stats().GetSdf()),
				dto.NewSprites(gb.ScissorsOfB().Sprites().Front())),
			dto.NewPokemon(
				gb.PaperOfB().ID(),
				gb.PaperOfB().Name(),
				string(gb.PaperOfB().Types().GetValue()),
				dto.NewStats(gb.PaperOfB().Stats().GetHP(),
					gb.PaperOfB().Stats().GetAtk(),
					gb.PaperOfB().Stats().GetDF(),
					gb.PaperOfB().Stats().GetSatk(),
					gb.PaperOfB().Stats().GetSdf()),
				dto.NewSprites(gb.PaperOfB().Sprites().Front())),
			dto.NewHP(gb.HPOfB().Current(), gb.HPOfB().Max()),
			dto.NewJaken(gb.JakenOfB().Name(), gb.JakenOfB().Value())),
		gb.Winner(),
	)
}

func GameBordFromDTO(gb dto.GameBord) E.Either[error, gamebord.GameBord] {
	return F.Pipe3(
		E.Right[error](gamebord.NewGameBord),
		E.Ap[func(partyB party.Party) func(winner O.Option[string]) gamebord.GameBord](partyFromDTO(gb.PartyA)),
		E.Ap[func(winner O.Option[string]) gamebord.GameBord](partyFromDTO(gb.PartyB)),
		E.Ap[gamebord.GameBord](E.Right[error](O.None[string]())),
	)
}

func pokemonFromDTO(poke dto.Pokemon) E.Either[error, pokemon.Pokemon] {
	return E.Map[error](
		func(typ types.Types) pokemon.Pokemon {
			return pokemon.New(poke.ID)(poke.Name)(typ)(stats.NewStats(poke.Stats.HP, poke.Stats.Atk, poke.Stats.Df, poke.Stats.SAtk, poke.Stats.SDf))(sprites.New(poke.Sprites.Front))
		},
	)(types.New(poke.Types))
}

func hpFromDTO(hitopoint dto.HP) hp.HP {
	return hp.New(hitopoint.Current)(hitopoint.Max)
}

func partyFromDTO(partyDTO dto.Party) E.Either[error, party.Party] {
	return F.Pipe4(
		E.Right[error](party.New(partyDTO.ID)),
		E.Ap[func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) func(hitpoint hp.HP) party.Party](pokemonFromDTO(partyDTO.Rock)),
		E.Ap[func(scissors pokemon.Pokemon) func(hitpoint hp.HP) party.Party](pokemonFromDTO(partyDTO.Paper)),
		E.Ap[func(hitpoint hp.HP) party.Party](pokemonFromDTO(partyDTO.Scissors)),
		E.Ap[party.Party](E.Right[error](hpFromDTO(partyDTO.HP))),
	)
}
