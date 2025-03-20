package sprites

type Sprites struct {
	front string
}

func New(front string) Sprites {
	return Sprites{
		front: front,
	}
}

func (s Sprites) Front() string {
	return s.front
}
