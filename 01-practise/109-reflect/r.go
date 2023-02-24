package _09_reflect

import (
	"golang.org/x/net/context"
)

type Person struct {
	name  string
	money uint64
}

func (p *Person) Say(ctx context.Context, str string) {

	p.money++
	//p.name = fmt.Sprintf("%s+++++", p.name)
}
func (p *Person) Bye() {
	//fmt.Println(p.name, "bye bye !")
	p.money++
}
