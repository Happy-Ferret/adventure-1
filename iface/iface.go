package iface

// Attack is an interface that can have one creature do damage to another
// creature.
type Attack interface {
	Name() string
	Description() string

	DoAttack(attacker, target Creature)
}

type Attacks []Attack

// Creature is an interface that represents something that can do attacks and
// has hit points.
type Creature interface {
	Name() string
	Description() string

	CurrentHitPoints() int
	Damage(damage int)
	IsDead() bool

	GetAttacks() Attacks
}

type BaseCreature struct {
	HitPoints int
}

func (c *BaseCreature) CurrentHitPoints() int {
	return c.HitPoints
}

func (c *BaseCreature) Damage(damage int) {
	c.HitPoints -= damage
	if c.HitPoints < 0 {
		c.HitPoints = 0
	}
}

func (c *BaseCreature) IsDead() bool {
	return c.HitPoints <= 0
}
