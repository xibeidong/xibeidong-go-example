package Singleton

import (
	"fmt"
	"testing"
)

func TestSingle(t *testing.T) {
	p := GetInstance()
	p.SetAge(3)
	fmt.Printf("%+v\n", *p)
}
