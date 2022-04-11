package main

type iSofa interface {
	setSofa(name string)
	setSize(size int)
	getSofa() string
	getSize() int
}
type sofa struct {
	name string
	size int
}

func (c *sofa) setSofa(name string) {
	c.name = name

}
func (c *sofa) setSize(size int) {
	c.size = size
}
func (c *sofa) getSofa() string {
	return c.name
}
func (c *sofa) getSize() int {
	return c.size
}
