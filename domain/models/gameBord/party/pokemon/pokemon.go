package pokemon

import (
	"my-go-app/domain/models/gameBord/party/pokemon/sprites"
	"my-go-app/domain/models/gameBord/party/pokemon/stats"
	"my-go-app/domain/models/gameBord/party/pokemon/types"
)

type Pokemon struct {
	id      int
	name    string
	types   types.Types
	stats   stats.Stats
	sprites sprites.Sprites
}

func New(id int) func(name string) func(typ types.Types) func(sts stats.Stats) func(sprt sprites.Sprites) Pokemon {
	return func(name string) func(typ types.Types) func(sts stats.Stats) func(sprt sprites.Sprites) Pokemon {
		return func(typ types.Types) func(sts stats.Stats) func(sprt sprites.Sprites) Pokemon {
			return func(sts stats.Stats) func(sprt sprites.Sprites) Pokemon {
				return func(sprt sprites.Sprites) Pokemon {
					return Pokemon{
						id:      id,
						name:    name,
						types:   typ,
						stats:   sts,
						sprites: sprt,
					}
				}
			}
		}
	}
}

// ID を取得
func (p Pokemon) ID() int {
	return p.id
}

// Name を取得
func (p Pokemon) Name() string {
	return p.name
}

// Types を取得
func (p Pokemon) Types() types.Types {
	return p.types
}

// Stats を取得
func (p Pokemon) Stats() stats.Stats {
	return p.stats
}

// Sprites を取得
func (p Pokemon) Sprites() sprites.Sprites {
	return p.sprites
}
