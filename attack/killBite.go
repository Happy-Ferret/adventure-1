package attack

import (
	"github.com/cosmodash/adventure/iface"
	"github.com/cosmodash/adventure/ui"
	"github.com/cosmodash/adventure/util"
)

type KillBite struct {
	power int
}

func NewKillBite(power int) iface.Attack {
	return &KillBite{power: power}
}

func (kB *KillBite) Name() string {
	return "Killing Bite"
}

func (kB *KillBite) Description() string {
	return "Bite the target on the neck"
}

func (kB *KillBite) DoAttack(attacker, target iface.Creature) {
	d := util.RandRange(1, 100)
	if d >= 90 {
		target.Damage(1000)
	}
	ui.ReportDamage(attacker, target, kB, d)
}
