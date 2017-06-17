package mob

import (
	"github.com/cosmodash/adventure/attack"
	"github.com/cosmodash/adventure/iface"
)

type KittyPet struct {
	iface.BaseCreature
}

func NewKittyPet() iface.Creature {
	return &KittyPet{
		iface.BaseCreature{HitPoints: 10},
	}
}

func (k *KittyPet) Name() string {
	return "Fluffy"
}

func (k *KittyPet) Description() string {
	return "A weak soft kitty pet."
}

func (k *KittyPet) GetAttacks() iface.Attacks {
	return iface.Attacks{attack.NewScratch(2)}
}
