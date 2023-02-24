package ProtoType

import (
	"fmt"
	"testing"
)

func TestCat_Clone(t *testing.T) {
	cat := Cat{
		name:   "cat1",
		weight: 20,
	}

	cat2 := cat.Clone()
	cat2.weight = 30
	cat2.name = "cat2"

	fmt.Printf("第一只猫叫%v,重 %v \n\n", cat.name, cat.weight)
	fmt.Printf("第二只猫叫%v,重 %v \n\n", cat2.name, cat2.weight)
}
