package main

type iChair interface {
	setChair(name string)
	setSize(size int)
	getChair() string
	getSize() int
}
type chair struct {
	name string
	size int
}

func (c *chair) setChair(name string) {
	c.name = name

}
func (c *chair) getChair() string {
	return c.name
}
func (c *chair) setSize(size int) {
	c.size = size
}

func (c *chair) getSize() int {
	return c.size
}
