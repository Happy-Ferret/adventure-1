package attack

import (
	"github.com/cosmodash/adventure/iface"
	"github.com/cosmodash/adventure/ui"
)

type Scratch struct {
	power int
}

func NewScratch(power int) iface.Attack {
	return &Scratch{power: power}
}

func (s *Scratch) Name() string {
	return "Scratch"
}

func (s *Scratch) Description() string {
	return "A quick swipe with claws."
}

func (s *Scratch) DoAttack(attacker, target iface.Creature) {
	target.Damage(s.power)
	ui.ReportDamage(attacker, target, s, s.power)
}
