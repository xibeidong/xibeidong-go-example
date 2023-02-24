package Facade

import "testing"

func TestNewCarFacade(t *testing.T) {
	f := NewCarFacade()
	f.CreateCompleteCar()
}
