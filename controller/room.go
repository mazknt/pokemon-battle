package controller

import (
	"log"
	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/model"
	"my-go-app/service"
)

type Room struct {
	room    connection.Room
	service service.Room
}

func NewRoomController(room connection.Room) *Room {
	return &Room{
		room:    room,
		service: service.NewRoomService(room),
	}
}

func (r *Room) Run() {
	for {
		select {
		case reqA := <-r.room.RegistRequest:
			reqB := <-r.room.RegistRequest
			log.Println("reqA: ", reqA)
			log.Println("reqB: ", reqB)
			r.regist(reqA, reqB)
		case battleField := <-r.room.Battle:
			reqA := <-r.room.JakenRequest
			reqB := <-r.room.JakenRequest
			r.jaken(reqA, reqB, battleField)
		}
	}
}

func (r *Room) regist(reqA dto.RegistRequestDto, reqB dto.RegistRequestDto) {
	battleField := r.service.GetPokemon(reqA, reqB)
	res := dto.ConvertRegistResponseDto(battleField)
	for _, client := range r.room.Clients {
		select {
		case client.RegistResponse <- res:
		default:
			// 送信に失敗
			close(client.RegistResponse)
		}
	}
	r.room.Battle <- battleField
}

func (r *Room) jaken(reqA dto.JakenRequestDto, reqB dto.JakenRequestDto, battleField model.BattleField) {
	res := r.service.Jaken(reqA, reqB, battleField)
	jekanResponse := dto.JakenResponseDto{
		PlayerA: res.PlayerA,
		PlayerB: res.PlayerB,
	}
	// すべてのクライアントにメッセージを送信
	for _, client := range r.room.Clients {
		select {
		case client.JakenResponse <- jekanResponse:
		default:
			// 送信に失敗
		}
	}
	r.room.Battle <- res
}
