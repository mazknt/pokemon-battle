package dto

import "my-go-app/model"

type RegistResponseDto struct {
	PlayerA model.User `json:"playerA"`
	PlayerB model.User `json:"playerB"`
}

type RegistRequestDto struct {
	Name     string `json:"name"`
	Paper    int    `json:"paper"`
	Rock     int    `json:"rock"`
	Scissors int    `json:"scissors"`
}

func ConvertRegistResponseDto(battleField model.BattleField) RegistResponseDto {
	return RegistResponseDto{
		PlayerA: battleField.PlayerA,
		PlayerB: battleField.PlayerB,
	}
}
