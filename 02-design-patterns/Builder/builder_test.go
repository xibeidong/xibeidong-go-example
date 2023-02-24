package Builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	b := NewConcreteBuilder1()
	director := NewDirector()
	director.BuildCar(b)
	car := b.GetProduct()
	fmt.Println(car)
}
