package game

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/cosmodash/adventure/hero"
	"github.com/cosmodash/adventure/iface"
	"github.com/cosmodash/adventure/mob"
	"github.com/cosmodash/adventure/ui"

	"github.com/fatih/color"
)

// Game keeps track of Creatures and runs the main game loop.
type Game struct {
	hero  iface.Creature
	enemy iface.Creature

	S *bufio.Scanner
}

func (g *Game) Run() {
	rand.Seed(time.Now().UnixNano())

	g.S = bufio.NewScanner(os.Stdin)

	heroName := ui.Prompt(g.S, "What is your name?")
	g.hero = hero.NewHero(heroName)

	color.New(color.FgHiWhite, color.Italic).Printf("Welcome %s! Prepare for battle!\n\n", g.hero.Name())

	g.enemy = g.createEnemy()
	for {
		g.runFightTurn()
		if g.enemy == nil {
			break
		}
		if g.hero.IsDead() {
			color.Red("You Lost! You are now dead. Sadness\n")
			break
		}
		if g.enemy.IsDead() {
			color.Green("You are victorious! You are a great and powerful Warrior!\n")
			break
		}
	}
}

func (g *Game) runFightTurn() {
	ui.PrintCreature("Hero", g.hero)
	ui.PrintCreature("Enemy", g.enemy)

	choice := ui.Menu(g.S, []string{"Attack", "Flee"})
	switch choice {
	case 0:
		g.runHeroAttack()
		if g.hero.IsDead() || g.enemy.IsDead() {
			return
		}
	case 1:
		color.Red("You flee the fight. Coward!\n")
		g.enemy = nil
		return
	}
	g.runEnemyAttack()
}

func (g *Game) runHeroAttack() {
	attacks := g.hero.GetAttacks()
	var attackStrings []string
	for _, a := range attacks {
		attackStrings = append(attackStrings, a.Name())
	}
	attackChoice := ui.Menu(g.S, attackStrings)
	attack := attacks[attackChoice]
	attack.DoAttack(g.hero, g.enemy)
}

func (g *Game) runEnemyAttack() {
	attacks := g.enemy.GetAttacks()
	attack := attacks[rand.Intn(len(attacks))]
	attack.DoAttack(g.enemy, g.hero)
}

func (g *Game) createEnemy() iface.Creature {
	return mob.NewKittyPet()
}
