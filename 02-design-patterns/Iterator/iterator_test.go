package Iterator

import (
	"fmt"
	"testing"
)

func TestConcreteIterator(t *testing.T) {
	data := ConcreteAggregate{array: []interface{}{}}
	iterator := data.GetIterator()
	iterator.Next()
	fmt.Println(iterator.Value())
}
