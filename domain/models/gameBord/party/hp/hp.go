package hp

import (
	"errors"
	"my-go-app/domain/models/gameBord/party/pokemon/stats"

	E "github.com/IBM/fp-go/either"
)

type HP struct {
	current int
	max     int
}

func Init(rock stats.Stats, scissors stats.Stats, paper stats.Stats) HP {
	return HP{
		current: rock.GetHP() + scissors.GetHP() + paper.GetHP(),
		max:     rock.GetHP() + scissors.GetHP() + paper.GetHP(),
	}
}

func New(current int) func(max int) HP {
	return func(max int) HP {
		return HP{
			current: current,
			max:     max,
		}
	}
}

func (hp HP) Current() int {
	return hp.current
}

func (hp HP) Max() int {
	return hp.max
}

func (hp HP) Damaged(damage int) E.Either[error, HP] {
	if hp.current-damage <= 0 {
		return E.Left[HP](errors.New("this party lose"))
	}
	return E.Right[error](HP{
		current: hp.current - damage,
		max:     hp.max,
	})
}
