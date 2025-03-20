package jakenDomainService

import (
	"errors"
	"my-go-app/domain/models/gameBord/jaken"

	E "github.com/IBM/fp-go/either"
)

func DetermineJakenWinnerDomainService(jakenA jaken.Jaken) func(jakenB jaken.Jaken) E.Either[error, jaken.Jaken] {
	return func(jakenB jaken.Jaken) E.Either[error, jaken.Jaken] {
		diff := (jakenA.Value() - jakenB.Value() + 3) % 3
		switch diff {
		case 1:
			return E.Right[error](jakenA)
		case 2:
			return E.Right[error](jakenB)
		default:
			return E.Left[jaken.Jaken](errors.New("draw"))
		}
	}
}
