package api

import (
	"errors"
	"fmt"
	"my-go-app/dto"
	"my-go-app/model/monad/result"
	"my-go-app/util"
	"net/http"
)

type PokeAPIInterface interface {
	GetPokemon(pokemonID int) result.Result[dto.PokemonResponseDto]
}

type PokeAPI struct{}

func (PokeAPIInterface PokeAPI) GetPokemon(pokemonID int) result.Result[dto.PokemonResponseDto] {
	// PokeAPIのURLを設定
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)

	// HTTP GETリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		return result.OfErr[dto.PokemonResponseDto](err)
	}
	defer resp.Body.Close()

	// ステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		return result.OfErr[dto.PokemonResponseDto](errors.New(fmt.Sprintf("failed to fetch data: %d", resp.StatusCode)))
	}

	// レスポンスボディを読み込み
	res := util.GetResponse[dto.PokemonResponseDto](resp)

	return res
}
