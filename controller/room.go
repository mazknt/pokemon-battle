package controller

import (
	"my-go-app/api"
	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/service"

	E "github.com/IBM/fp-go/either"
)

type Room struct {
	room    connection.Room
	service service.GameBordService
}

func NewRoomController(room connection.Room) *Room {
	return &Room{
		room:    room,
		service: service.NewGameBord(api.NewPokeAPI()),
	}
}

func (r *Room) Run() {
	for {
		select {
		case reqA := <-r.room.RegistRequest:
			reqB := <-r.room.RegistRequest
			r.regist(reqA, reqB)
		case battleFieldE := <-r.room.Battle:
			reqA := <-r.room.JakenRequest
			reqB := <-r.room.JakenRequest
			E.Fold(
				func(err error) string { return "" },
				func(battleField dto.GameBord) string {
					r.jaken(reqA, reqB, battleField)
					return ""
				},
			)(battleFieldE)

		}
	}
}

func (r *Room) regist(reqA dto.RegistRequestDto, reqB dto.RegistRequestDto) {
	battleField := r.service.CreateGameBord(reqA)(reqB)
	for _, client := range r.room.Clients {
		select {
		case client.RegistResponse <- battleField:
		default:
			close(client.RegistResponse)
		}
	}
	r.room.Battle <- battleField
}

func (r *Room) jaken(reqA dto.JakenRequestDto, reqB dto.JakenRequestDto, battleField dto.GameBord) {
	res := r.service.Battle(battleField)(reqA)(reqB)
	for _, client := range r.room.Clients {
		select {
		case client.JakenResponse <- res:
		default:
			// 送信に失敗
		}
	}
	r.room.Battle <- res
}
