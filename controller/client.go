package controller

import (
	"encoding/json"
	"log"

	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/model/monad/result"
	"my-go-app/util"

	"github.com/gorilla/websocket"
)

type clientController struct {
	client *connection.Client
}

func (c *clientController) Read(manager *connection.Manager) {
	room := manager.Room[c.client.RoomID]
	for {
		_, msg, err := c.client.Socket.ReadMessage()
		if err != nil {
			break
		}
		res := util.JsonUnmarshal[dto.RequestTest](msg)
		req := result.Unwrap[dto.RequestTest](res)
		if util.IsUndefined[dto.RequestTest](req) {
			break
		}
		// room情報の更新
		select {
		case roomInfo := <-c.client.RoomInfo:
			room = roomInfo
		default:
		}

		// リクエストの処理
		switch req.URL {
		case "connect":
			c.client.ID = util.JsonDecode[dto.ConnectRequest](req).ID
			manager.ConnectRequest <- c.client
		case "getPokemon":
			room.JakenRequest <- util.JsonDecode[dto.JakenRequestDto](req)
		case "regist":
			// RoomをRun()させる必要があり、それはmanager.Run()内でするべき
			manager.RunRequest <- room
			room.RegistRequest <- util.JsonDecode[dto.RegistRequestDto](req)
		case "send-battle-request":
			manager.BattleRequestRequest <- util.JsonDecode[dto.BattleRequestRequest](req)
		}
	}
}

func (c *clientController) Write() {
	var err error
	for {
		select {
		case res := <-c.client.RegistResponse:
			response, _ := json.Marshal(res)
			err = c.client.Socket.WriteMessage(websocket.TextMessage, response)
		case res := <-c.client.JakenResponse:
			response, _ := json.Marshal(res)
			err = c.client.Socket.WriteMessage(websocket.TextMessage, response)
		case res := <-c.client.BattleRequestFromOther:
			response, _ := json.Marshal(res)
			err = c.client.Socket.WriteMessage(websocket.TextMessage, response)
		case res := <-c.client.BattleRequestResponse:
			response, _ := json.Marshal(res)
			err = c.client.Socket.WriteMessage(websocket.TextMessage, response)
		}
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
	c.client.Socket.Close()
}
