package ui

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/cosmodash/adventure/iface"
	"github.com/fatih/color"
)

func Prompt(s *bufio.Scanner, prompt string) string {
	p := color.New(color.FgMagenta)
	p.Printf("%s ", prompt)
	s.Scan()
	return s.Text()
}

func Menu(s *bufio.Scanner, options []string) int {
	n := color.New(color.FgWhite, color.Underline)
	d := color.New(color.FgWhite)
	e := color.New(color.FgRed)
	for {
		for i, o := range options {
			n.Printf("%d", i+1)
			d.Printf(" %s\n", o)
		}

		choiceString := Prompt(s, "What do you want to do?")
		choice, _ := strconv.Atoi(choiceString)
		if choice <= 0 || choice > len(options) {
			e.Printf("Please pick one of the options.\n")
		} else {
			return choice - 1
		}
	}
}

func ReportDamage(attacker, target iface.Creature, attack iface.Attack, damage int) {
	d := color.New(color.BgRed, color.FgHiWhite)
	if damage > 0 {
		d.Printf("%s uses %s on %s and does %d damage.\n",
			attacker.Name(), target.Name(), attack.Name(), damage)
	} else {
		d.Printf("%s uses %s on %s but misses!\n",
			attacker.Name(), target.Name(), attack.Name())
	}

	if attacker.IsDead() {
		d.Printf("%s is dead!\n", attacker.Name())
	}
	if target.IsDead() {
		d.Printf("%s is dead!\n", target.Name())
	}
}

func PrintCreature(header string, c iface.Creature) {
	h := color.New(color.Underline, color.FgHiWhite)
	n := color.New(color.FgBlue, color.Bold)
	k := color.New(color.FgYellow, color.Bold)
	v := color.New(color.FgWhite)
	h.Println(header)
	n.Printf("%s: ", c.Name())
	v.Printf("%s\n", c.Description())
	k.Printf("Hit Points: ")
	v.Printf("%d\n", c.CurrentHitPoints())

	k.Printf("Attacks:\n")
	for _, a := range c.GetAttacks() {
		k.Printf("  %s: ", a.Name())
		v.Printf("%s\n", a.Description())
	}

	fmt.Println()
}
