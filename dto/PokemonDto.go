package dto

import "my-go-app/model"

// PokemonResponse ポケモンの詳細情報を持つ構造体
type PokemonResponseDto struct {
	ID      int                                 `json:"id"`
	Name    string                              `json:"name"`
	Types   []struct{ Type model.PokemonTypes } `json:"types"`
	Stats   []model.Stat                        `json:"stats"`
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
