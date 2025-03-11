package connection

import (
	"my-go-app/dto"
)

type Manager struct {
	Room                 map[string]*Room
	ConnectRequest       chan *Client
	Client               map[string]*Client
	BattleRequestRequest chan dto.BattleRequestRequest
	WaitingClients       map[Client](string) // 誰が誰を待っているかわかるように
	RunRequest           chan *Room
}

func NewManager() *Manager {
	return &Manager{
		Room:                 make(map[string]*Room),
		ConnectRequest:       make(chan *Client),
		Client:               make(map[string]*Client),
		BattleRequestRequest: make(chan dto.BattleRequestRequest),
		WaitingClients:       make(map[Client](string)),
		RunRequest:           make(chan *Room, 1),
	}
}
