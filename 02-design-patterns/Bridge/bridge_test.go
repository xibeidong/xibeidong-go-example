package Bridge

import "testing"

func TestCoffee(t *testing.T) {
	coffee := LargeCoffee{}
	coffee.additives = &Milk{}

	coffee.Volume()
	coffee.additives.AddSomething()
}
