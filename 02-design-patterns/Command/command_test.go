package Command

import (
	"testing"
)

func TestCommand(t *testing.T) {
	invoker := Invoker{}
	invoker.addCommand(&ConcreteCommand1{})
	invoker.addCommand(&ConcreteCommand2{})

	invoker.Call()
}
