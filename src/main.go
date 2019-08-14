package main

import "fmt"

type parents struct {
	model string
}

func (p *parents) Init(m string) {
	p.model = m
}
func (p *parents) Get() {
	fmt.Println(p.model)
}

type child struct {
	parents
	model int
}

func (ch *child) Init(i int) {
	ch.model = i
}

func main() {
	ch := new(child)
	ch.parents.Init("Ljan")
	ch.Get()
}
