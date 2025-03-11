package dto

type BattleRequestRequest struct {
	To   string `json:"to"`
	From string `json:"from"`
}

type BattleRequestResponse struct {
	To        string `json:"playerA"`
	From      string `json:"playerB"`
	IsMatched bool   `json:"isMatched"`
}
