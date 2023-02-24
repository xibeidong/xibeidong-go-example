package Builder

import (
	"fmt"
	"testing"
)

func TestCar(t *testing.T) {
	carBuilder := CarBuilder{}
	car := carBuilder.
		EngineBuild("engine").
		RackBuild("rack").
		SteeringBuild("steering").
		Build()
	fmt.Println(car)
}
