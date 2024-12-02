package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	a := true
	if firstAttacker == fighter1.Name {
		for a {
			fighter2.Health = fighter2.Health - fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
			fighter1.Health = fighter1.Health - fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
		}
	} else {
		for a {
			fighter1.Health = fighter1.Health - fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
			fighter2.Health = fighter2.Health - fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
		}
	}
	return "First Attacker is missing"
}

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))
}
