package Factory

import "fmt"

//	工厂模式

type Monster interface {
	Attack()
	Die()
}

// Slime 史莱姆
type Slime struct {
}

func (slime *Slime) Die() {
	fmt.Println("slime is dead!!!")
}

func (slime *Slime) Attack() {
	fmt.Println("slime attack once!")
}

// Werewolf 狼人
type Werewolf struct {
}

func (wolf *Werewolf) Die() {
	fmt.Println("wolf is dead!!!")
}

func (wolf *Werewolf) Attack() {
	fmt.Println("wolf attack once!")
}

func NewMonster(name string) Monster {
	switch name {
	case "wolf":
		return &Werewolf{}
	case "slime":
		return &Slime{}
	default:
		return &Slime{}

	}
}
