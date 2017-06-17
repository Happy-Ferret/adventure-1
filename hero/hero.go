package hero

import "github.com/cosmodash/adventure/iface"
import "github.com/cosmodash/adventure/attack"

type Hero struct {
	iface.BaseCreature
	HeroName string
}

func NewHero(name string) iface.Creature {
	return &Hero{
		BaseCreature: iface.BaseCreature{HitPoints: 20},
		HeroName:     name,
	}
}

func (h *Hero) Name() string {
	return h.HeroName
}

func (h *Hero) Description() string {
	return "A fine looking Warrior cat"
}

func (h *Hero) GetAttacks() iface.Attacks {
	return iface.Attacks{
		attack.NewPounce(4),
	}
}
