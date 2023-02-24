package Composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	bigBag := Bags{
		name: "big",
		bags: nil,
	}
	bigBag.Add(&Goods{
		name:      "西红柿",
		unitPrice: 5,
		quantity:  10,
	})

	smallBag := Bags{
		name: "small",
		bags: nil,
	}
	smallBag.Add(&Goods{
		name:      "黄瓜",
		unitPrice: 2,
		quantity:  3,
	})
	bigBag.Add(&smallBag)
	p := bigBag.Calculation()
	fmt.Println(p)
}
