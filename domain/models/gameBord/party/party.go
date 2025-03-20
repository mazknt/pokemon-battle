package party

import (
	"errors"
	"my-go-app/domain/models/gameBord/jaken"
	"my-go-app/domain/models/gameBord/party/hp"
	"my-go-app/domain/models/gameBord/party/pokemon"

	E "github.com/IBM/fp-go/either"
	F "github.com/IBM/fp-go/function"
)

type Party struct {
	id       string
	rock     pokemon.Pokemon
	scissors pokemon.Pokemon
	paper    pokemon.Pokemon
	hp       hp.HP
	jaken    jaken.Jaken
}

func Init(id string) func(rock pokemon.Pokemon) func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) Party {
	return func(rock pokemon.Pokemon) func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) Party {
		return func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) Party {
			return func(scissors pokemon.Pokemon) Party {
				return Party{
					id:       id,
					rock:     rock,
					paper:    paper,
					scissors: scissors,
					hp:       hp.Init(rock.Stats(), paper.Stats(), scissors.Stats()),
					jaken:    jaken.NewRock(), // デフォルト値
				}
			}
		}
	}
}

func New(id string) func(rock pokemon.Pokemon) func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) func(hitpoint hp.HP) Party {
	return func(rock pokemon.Pokemon) func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) func(hitpoint hp.HP) Party {
		return func(paper pokemon.Pokemon) func(scissors pokemon.Pokemon) func(hitpoint hp.HP) Party {
			return func(scissors pokemon.Pokemon) func(hitpoint hp.HP) Party {
				return func(hitpoint hp.HP) Party {
					return Party{
						id:       id,
						rock:     rock,
						paper:    paper,
						scissors: scissors,
						hp:       hitpoint,
						jaken:    jaken.NewRock(), // デフォルト値
					}
				}
			}
		}
	}

}

func (p Party) ID() string {
	return p.id
}

func (p Party) Rock() pokemon.Pokemon {
	return p.rock
}

func (p Party) Scissors() pokemon.Pokemon {
	return p.scissors
}

func (p Party) Paper() pokemon.Pokemon {
	return p.paper
}

func (p Party) HP() hp.HP {
	return p.hp
}

func (p Party) Jaken() jaken.Jaken {
	return p.jaken
}

func (p Party) Damaged(damage int) E.Either[error, Party] {
	return F.Pipe2(
		p.hp.Damaged(damage),
		E.Map[error](
			func(h hp.HP) Party {
				return Party{
					id:       p.id,
					rock:     p.rock,
					scissors: p.scissors,
					paper:    p.paper,
					hp:       h,
				}
			}),
		E.Fold(
			func(err error) E.Either[error, Party] {
				if err.Error() == "this party lose" {
					E.Left[Party](errors.New(p.ID()))
				}
				return E.Left[Party](err)
			},
			func(party Party) E.Either[error, Party] { return E.Right[error](party) },
		),
	)
}
