package api

import (
	"errors"
	"fmt"
	"my-go-app/dto"
	"my-go-app/util"
	"net/http"

	E "github.com/IBM/fp-go/either"
)

type PokeAPI struct{}

func NewPokeAPI() PokeAPI {
	return PokeAPI{}
}

func (pokeAPI PokeAPI) GetPokemon(id int) E.Either[error, dto.PokemonResponseDto] {
	// PokeAPIのURLを設定
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id)

	// HTTP GETリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		return E.Left[dto.PokemonResponseDto](err)
	}
	defer resp.Body.Close()

	// ステータスコードをチェック
	if resp.StatusCode != http.StatusOK {
		return E.Left[dto.PokemonResponseDto](errors.New(fmt.Sprintf("failed to fetch data: %d", resp.StatusCode)))
	}

	// レスポンスボディを読み込み
	return util.GetResponse[dto.PokemonResponseDto](resp)

}
