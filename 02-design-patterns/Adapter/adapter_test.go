package Adapter

import "testing"

func TestAdapter(t *testing.T) {
	var motor Motor
	motor = &ElectricAdapter{elect: ElectricEngine{}}
	motor.Driver()

	motor = &OilAdapter{oil: OilEngine{}}
	motor.Driver()
}
