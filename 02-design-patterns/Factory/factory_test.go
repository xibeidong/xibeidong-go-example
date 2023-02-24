package Factory

import "testing"

func TestNewMonster(t *testing.T) {
	slime := NewMonster("slime")
	slime.Attack()
	slime.Die()
	wolf := NewMonster("wolf")
	wolf.Attack()
	wolf.Die()
}
