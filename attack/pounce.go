package attack

import (
	"github.com/cosmodash/adventure/iface"
	"github.com/cosmodash/adventure/ui"
)

type Pounce struct {
	power int
}

func NewPounce(power int) iface.Attack {
	return &Pounce{power: power}
}

func (p *Pounce) Name() string {
	return "Pounce"
}

func (p *Pounce) Description() string {
	return "Jump on the other cat and wrestle it to the ground."
}

func (p *Pounce) DoAttack(attacker, target iface.Creature) {
	target.Damage(p.power)
	ui.ReportDamage(attacker, target, p, p.power)
}
