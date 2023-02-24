package Strategy

import "testing"

func TestChickenCook(t *testing.T) {
	kitchen := &Kitchen{}
	kitchen.setStrategy(&chickenRoast{})
	kitchen.Cook()
	kitchen.setStrategy(&chickenSoup{})
	kitchen.Cook()
}
