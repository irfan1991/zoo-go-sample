package tiger

import "gitlab.com/ari-salt/zoo/animal/cat"

type tiger struct {
	cat.Cat
}

func New(name, kind string) *tiger {
	c := cat.New(name, kind)

	return &tiger{
		Cat: *c,
	}
}

func (t *tiger) Rawr() string {
	return "Rawr!"
}
