package stats

type Stats struct {
	hp   int
	atk  int
	df   int
	satk int
	sdf  int
}

func NewStats(hp, atk, df, satk, sdf int) Stats {
	return Stats{
		hp:   hp,
		atk:  atk,
		df:   df,
		satk: satk,
		sdf:  sdf,
	}
}

func (s Stats) GetHP() int {
	return s.hp
}

func (s Stats) GetAtk() int {
	return s.atk
}

func (s Stats) GetDF() int {
	return s.df
}

func (s Stats) GetSatk() int {
	return s.satk
}

func (s Stats) GetSdf() int {
	return s.sdf
}
