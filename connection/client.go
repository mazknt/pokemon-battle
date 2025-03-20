package connection

import (
	"my-go-app/dto"
	"my-go-app/model"

	E "github.com/IBM/fp-go/either"
	"github.com/gorilla/websocket"
)

type Client struct {
	// クライアントの接続
	Socket *websocket.Conn
	ID     string

	RegistRequest  chan dto.RegistRequestDto
	RegistResponse chan E.Either[error, dto.GameBord]

	JakenRequest  chan dto.JakenRequestDto
	JakenResponse chan E.Either[error, dto.GameBord]

	BattleRequestRequest  chan dto.BattleRequestRequest
	BattleRequestResponse chan dto.BattleRequestResponse

	BattleRequestFromOther chan dto.BattleRequestRequest // 他人からの果たし状をやり取り

	RoomInfo chan *Room
	Battle   chan model.BattleField
	RoomID   string
}

func NewClient(socket *websocket.Conn, roomID string) *Client {
	return &Client{
		Socket: socket,
		ID:     "",

		RegistRequest:  make(chan dto.RegistRequestDto),
		RegistResponse: make(chan E.Either[error, dto.GameBord]),

		JakenRequest:  make(chan dto.JakenRequestDto),
		JakenResponse: make(chan E.Either[error, dto.GameBord]),

		BattleRequestRequest:   make(chan dto.BattleRequestRequest),
		BattleRequestResponse:  make(chan dto.BattleRequestResponse),
		BattleRequestFromOther: make(chan dto.BattleRequestRequest),

		RoomInfo: make(chan *Room, 2),
		Battle:   make(chan model.BattleField),
		RoomID:   roomID,
	}
}
