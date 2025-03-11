package model

type User struct {
	Name  string `json:"name"`
	Party Party  `json:"party"`
}

// Party 構造体の定義
type Party struct {
	Rock     Pokemon `json:"rock"`
	Paper    Pokemon `json:"paper"`
	Scissors Pokemon `json:"scissors"`
	HP       HP      `json:"hp"`
}

// HP 構造体の定義
type HP struct {
	Max     int `json:"max"`
	Current int `json:"current"`
}
