package model

// PokemonType はポケモンのタイプを表す型
type PokemonType string

const (
	Normal   PokemonType = "normal"
	Fighting PokemonType = "fighting"
	Flying   PokemonType = "flying"
	Poison   PokemonType = "poison"
	Ground   PokemonType = "ground"
	Rock     PokemonType = "rock"
	Bug      PokemonType = "bug"
	Ghost    PokemonType = "ghost"
	Steel    PokemonType = "steel"
	Fire     PokemonType = "fire"
	Water    PokemonType = "water"
	Grass    PokemonType = "grass"
	Electric PokemonType = "electric"
	Psychic  PokemonType = "psychic"
	Ice      PokemonType = "ice"
	Dragon   PokemonType = "dragon"
	Dark     PokemonType = "dark"
	Fairy    PokemonType = "fairy"
)

// PokemonTypes ポケモンのタイプを持つ構造体
type PokemonTypes struct {
	Name PokemonType `json:"name"`
}

// Stat ポケモンのステータス情報を持つ構造体
type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type Pokemon struct {
	ID      int                           `json:"id"`
	Name    string                        `json:"name"`
	Types   []struct{ Type PokemonTypes } `json:"types"`
	Stats   []Stat                        `json:"stats"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Species struct {
		FlavorTextEntries []struct {
			FlavorText string `json:"flavor_text"`
			Language   struct {
				Name string `json:"name"`
			} `json:"language"`
		} `json:"flavor_text_entries"`
	} `json:"species"`
}
