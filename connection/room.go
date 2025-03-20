package connection

import (
	"my-go-app/dto"

	E "github.com/IBM/fp-go/either"
)

type Room struct {
	// 登録情報のやり取り
	RegistRequest  chan dto.RegistRequestDto
	RegistResponse chan dto.RegistResponseDto
	JakenRequest   chan dto.JakenRequestDto
	JakenResponse  chan dto.JakenResponseDto
	Battle         chan E.Either[error, dto.GameBord]
	Clients        []Client
	Manager        Manager
	ID             string
}

func NewRoom(client Client, id string) *Room {
	return &Room{
		ID:             id,
		RegistRequest:  make(chan dto.RegistRequestDto),
		RegistResponse: make(chan dto.RegistResponseDto),
		Battle:         make(chan E.Either[error, dto.GameBord], 1),
		JakenRequest:   make(chan dto.JakenRequestDto),
		JakenResponse:  make(chan dto.JakenResponseDto),
	}
}
