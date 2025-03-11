package connection

import (
	"my-go-app/dto"
	"my-go-app/model"
)

type Room struct {
	// 登録情報のやり取り
	RegistRequest  chan dto.RegistRequestDto
	RegistResponse chan dto.RegistResponseDto
	JakenRequest   chan dto.JakenRequestDto
	JakenResponse  chan dto.JakenResponseDto
	Battle         chan model.BattleField
	Clients        []Client
	Manager        Manager
	ID             string
}

func NewRoom(client Client, id string) *Room {
	return &Room{
		ID:             id,
		RegistRequest:  make(chan dto.RegistRequestDto),
		RegistResponse: make(chan dto.RegistResponseDto),
		Battle:         make(chan model.BattleField, 1),
		JakenRequest:   make(chan dto.JakenRequestDto),
		JakenResponse:  make(chan dto.JakenResponseDto),
	}
}
