package dto

import "my-go-app/model"

type JakenRequestDto struct {
	Name  string      `json:"name"`
	Jaken model.Jaken `json:"jaken"`
}
type JakenResponseDto struct {
	PlayerA model.User `json:"playerA"`
	PlayerB model.User `json:"playerB"`
}

type RequestTest struct {
	URL  string `json:"url"`
	Json string `json:"json"`
}
