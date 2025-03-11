package main

import (
	"fmt"
	"log"
	"my-go-app/connection"
	"my-go-app/controller"
	"net/http"
)

func main() {
	// WebSocket接続を受け付けるエンドポイント
	managerController := controller.NewManagerController(connection.NewManager())
	go managerController.Run()
	// go room.Jaken()
	http.Handle("/get-pokemons", managerController)

	// サーバーをポート8090で起動
	fmt.Println("WebSocket server is running on ws://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
