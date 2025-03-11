package service

import (
	"my-go-app/connection"
	"my-go-app/dto"
	"my-go-app/util"
	"sort"

	Arr "github.com/IBM/fp-go/array"
	FP "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
)

type ManagerService struct {
	manager *connection.Manager
}

func NewManagerService(manager *connection.Manager) *ManagerService {
	return &ManagerService{
		manager: manager,
	}
}

func (m ManagerService) MakeBattleRoom(battleRequest dto.BattleRequestRequest) {
	waitingClients := util.GetMapKeys(m.manager.WaitingClients)
	waitingClient := *m.manager.Client[battleRequest.From]
	newRoom := FP.Pipe1(
		getTargetClientOption(battleRequest.To, waitingClients),
		O.Fold(
			// ない場合
			func() connection.Room {
				// clientを待機させる
				newRoom := makeNewBattleRoom(waitingClient, battleRequest.To)
				waitingClient.RoomID = newRoom.ID
				m.manager.WaitingClients[waitingClient] = battleRequest.To
				return newRoom
			},
			// ある場合
			func(targetClient connection.Client) connection.Room {
				// 相手の待ち人が自分ではない場合、新しく部屋を作る。
				if m.manager.WaitingClients[targetClient] != battleRequest.From {
					m.manager.WaitingClients[waitingClient] = battleRequest.To
					return makeNewBattleRoom(waitingClient, battleRequest.To)
				}
				// 相手の部屋に入る
				targetRoom := m.manager.Room[targetClient.RoomID]
				pushClient := m.manager.Client[battleRequest.From]
				delete(m.manager.WaitingClients, targetClient)
				return updateBattleRoom(*targetRoom)(*pushClient)
			},
		),
	)
	m.manager.Room[newRoom.ID] = &newRoom
	for _, client := range m.manager.Room[newRoom.ID].Clients {
		client.RoomInfo <- &newRoom
	}
}

func getClientWithId(clients []connection.Client) func(string) connection.Client {
	return func(clientID string) connection.Client {
		return FP.Pipe3(clients,
			Arr.Filter(func(client connection.Client) bool { return clientID == client.ID }),
			util.IsExist,
			util.Undefined,
		)
	}
}

func makeNewBattleRoom(waitingClient connection.Client, targetClientID string) connection.Room {
	roomID := makeRoomID(waitingClient.ID, targetClientID)
	newRoom := connection.NewRoom(waitingClient, roomID)
	newRoom.Clients = append(newRoom.Clients, waitingClient)
	return *newRoom
}

func makeRoomID(str1, str2 string) string {
	// 文字列をスライスに入れてソート
	strs := []string{str1, str2}
	sort.Strings(strs)
	// ソート後の文字列を結合
	return strs[0] + strs[1]
}

func updateBattleRoom(targetRoom connection.Room) func(pushClient connection.Client) connection.Room {
	return func(pushClient connection.Client) connection.Room {
		clients := util.Push(targetRoom.Clients)(pushClient)
		targetRoom.Clients = clients
		return targetRoom
	}
}

func UpdateRooms(targetClient connection.Client, rooms []connection.Room, pushClient connection.Client) []connection.Room {
	return FP.Pipe1(
		getTargetRoomOption(targetClient, rooms),
		O.Fold(
			func() []connection.Room { return rooms },
			func(targetRoom connection.Room) []connection.Room {
				return FP.Pipe2(
					targetRoom,
					func(room connection.Room) connection.Room {
						room.Clients = append(room.Clients, pushClient)
						return room
					},
					util.Update(rooms, targetRoom),
				)
			},
		),
	)
}

func getTargetRoomOption(targetClient connection.Client, rooms []connection.Room) O.Option[connection.Room] {
	return FP.Pipe2(
		rooms,
		Arr.Filter(func(room connection.Room) bool {
			return FP.Pipe1(
				getTargetClientOption(targetClient.ID, room.Clients),
				util.IsUndefined,
			)
		}),
		util.IsExist,
	)
}

func getTargetClientOption(targetClientId string, clients []connection.Client) O.Option[connection.Client] {
	return FP.Pipe2(
		clients,
		Arr.Filter(func(client connection.Client) bool {
			return client.ID == targetClientId
		}),
		util.IsExist,
	)
}
