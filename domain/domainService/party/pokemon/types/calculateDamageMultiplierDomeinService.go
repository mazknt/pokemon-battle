package typesDomainService

import (
	"my-go-app/constants"
	"my-go-app/domain/models/gameBord/party/pokemon/types"
)

func CalculateDamageMultiplier(winner types.Types, looser types.Types) float32 {
	return constants.TypeEffectiveness[string(winner.GetValue())][string(looser.GetValue())]
}
