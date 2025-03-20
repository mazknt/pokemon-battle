package pokemonDomainService

import (
	typesDomainService "my-go-app/domain/domainService/party/pokemon/types"
	"my-go-app/domain/models/gameBord/party/pokemon"
)

func CalculateDamage(winner pokemon.Pokemon) func(loser pokemon.Pokemon) int {
	return func(loser pokemon.Pokemon) int {
		attackDamage :=
			(22*75*winner.Stats().GetAtk())/loser.Stats().GetDF()/50 +
				2
		specialAttackDamage :=
			(22*75*winner.Stats().GetSatk())/loser.Stats().GetSdf()/50 +
				2

		damageRate := typesDomainService.CalculateDamageMultiplier(winner.Types(), loser.Types())
		if attackDamage >= specialAttackDamage {
			return int(float32(attackDamage) * damageRate)
		}
		return int(float32(specialAttackDamage) * damageRate)
	}
}
