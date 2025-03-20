package responsedto

import "my-go-app/dto"

type Jaken struct {
	PlayerA dto.Party `json:"playerA"`
	PlayerB dto.Party `json:"playerB"`
}

func NewJaken(playerA dto.Party) func(playerB dto.Party) Jaken
