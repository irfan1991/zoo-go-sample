package cat

type Cat struct {
	name string
	kind string
}

func New(name, kind string) *Cat {
	return &Cat{
		name: name,
		kind: kind,
	}
}

func (c *Cat) Name() string {
	return c.name
}

func (c *Cat) Kind() string {
	return c.kind
}

func (c *Cat) Rawr() string {
	return "Meow..."
}
