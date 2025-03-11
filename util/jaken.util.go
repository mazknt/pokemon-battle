package util

import (
	"log"
	"my-go-app/constants"
	"my-go-app/dto"
	"my-go-app/model"
)

func JudgeJaken(playerA dto.JakenRequestDto, playerB dto.JakenRequestDto) dto.JakenRequestDto {
	log.Println(playerA)
	log.Println(playerB)
	switch playerA.Jaken.Value - playerB.Jaken.Value {
	case -1:
		return playerA
	case 2:
		return playerA
	case -2:
		return playerA
	case 1:
		return playerB
	default:
		var zeroValue dto.JakenRequestDto
		return zeroValue
	}
}

func CalculateDamage(battleField model.BattleField, winner dto.JakenRequestDto) model.BattleField {
	if winner.Name == "" {
		return battleField
	}
	battlePokemons := getAttackerDifender(battleField, winner)
	damage := calculateDamage(battlePokemons.Attacker, battlePokemons.Difender)
	return updateBattleField(battleField, winner, damage)
}

func getAttackerDifender(battleField model.BattleField, winner dto.JakenRequestDto) model.BattlePokemon {
	if winner.Name == battleField.PlayerA.Name {
		return model.BattlePokemon{
			Attacker: getBattlePokemon(battleField.PlayerA.Party, winner.Jaken),
			Difender: getBattlePokemon(battleField.PlayerB.Party, winner.Jaken),
		}
	}
	return model.BattlePokemon{
		Difender: getBattlePokemon(battleField.PlayerA.Party, winner.Jaken),
		Attacker: getBattlePokemon(battleField.PlayerB.Party, winner.Jaken),
	}
}

// winnerのじゃんけんの手が何かを特定し、そのじゃんけんのポケモンを取得する
func getBattlePokemon(party model.Party, jaken model.Jaken) model.Pokemon {
	switch jaken.Name {
	case "rock":
		return party.Rock
	case "paper":
		return party.Paper
	default:
		return party.Scissors
	}
}

func calculateDamageRate(attacker model.Pokemon, difender model.Pokemon) float32 {
	return constants.TypeEffectiveness[string(attacker.Types[0].Type.Name)][string(difender.Types[0].Type.Name)]
}

func calculateDamage(attacker model.Pokemon, difender model.Pokemon) int {
	damageRate := calculateDamageRate(attacker, difender)
	attackDamage :=
		(22*75*attacker.Stats[1].BaseStat)/difender.Stats[1].BaseStat/50 +
			2
	specialAttackDamage :=
		(22*75*attacker.Stats[3].BaseStat)/difender.Stats[3].BaseStat/50 +
			2
	if attackDamage >= specialAttackDamage {
		return int(float32(attackDamage) * damageRate)
	}
	return int(float32(specialAttackDamage) * damageRate)
}

func updateBattleField(battleField model.BattleField, winner dto.JakenRequestDto, damage int) model.BattleField {
	if winner.Name == battleField.PlayerA.Name {
		battleField.PlayerB.Party.HP.Current -= damage
	} else {
		battleField.PlayerA.Party.HP.Current -= damage
	}
	return battleField
}
