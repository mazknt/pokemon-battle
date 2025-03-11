package model

// BattleField型の定義
type BattleField struct {
	PlayerA User `json:"playerA"`
	PlayerB User `json:"playerB"`
}

type BattlePokemon struct {
	Attacker Pokemon
	Difender Pokemon
}
