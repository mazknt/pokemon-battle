package controller

import (
	"log"
	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/service"
	"net/http"

	"github.com/gorilla/websocket"
)

type ManagerController struct {
	manager *connection.Manager
	service *service.ManagerService
}

func (m ManagerController) Run() {
	for {
		select {
		case client := <-m.manager.ConnectRequest:
			m.manager.Client[client.ID] = client
		case request := <-m.manager.BattleRequestRequest:
			beforeLength := len(m.manager.Room)
			m.service.MakeBattleRoom(request)
			afterLength := len(m.manager.Room)
			// 部屋を新規に作成した場合、相手に果たし状を送る
			if beforeLength != afterLength {
				// 相手がログインしていないと↓のコードは破綻する
				m.manager.Client[request.To].BattleRequestFromOther <- request
			} else {
				// 部屋に人が揃った場合、バトルのセットを送る
				m.manager.Client[request.To].BattleRequestResponse <- dto.BattleRequestResponse{
					From:      request.To,
					To:        request.From,
					IsMatched: true,
				}
				m.manager.Client[request.From].BattleRequestResponse <- dto.BattleRequestResponse{
					From:      request.From,
					To:        request.To,
					IsMatched: true,
				}
			}
		case runReuquest := <-m.manager.RunRequest:
			roomController := NewRoomController(*m.manager.Room[runReuquest.ID])
			go roomController.Run()
		}
	}
}

func NewManagerController(manager *connection.Manager) *ManagerController {
	return &ManagerController{
		manager: manager,
		service: service.NewManagerService(manager),
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 任意のOriginからの接続を許可（セキュリティ上の理由で実際のアプリではチェックを行うべき）
		return true
	},
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

func (m *ManagerController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := connection.NewClient(socket, "")
	controller := clientController{
		client: client,
	}
	go controller.Read(m.manager)
	controller.Write()
}
