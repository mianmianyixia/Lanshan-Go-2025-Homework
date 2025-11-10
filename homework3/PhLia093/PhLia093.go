package PhLia093

import (
	"fmt"
)

type Character struct {
	Attacks float64
	Defence float64
	Hp      float64
	IsAlive bool
}
type Action interface {
	Attack(*Character)
	Defend(*Character)
	Skill(*Character)
	Hurt(*Character)
}

func (xilian *Character) Attack(enemy *Character) { //攻击
	(*enemy).Hp -= xilian.Attacks
	if (*enemy).Hp <= 0 {
		(*enemy).IsAlive = false
		fmt.Println("击败")
	}
}
func (xilian *Character) Defend(enemy *Character) { //防御
	if (*xilian).Defence-(*enemy).Attacks >= 0 {
		fmt.Println("未造成任何伤害")
		(*xilian).Hp += (*enemy).Attacks
	} else {
		(*xilian).Hp = (*xilian).Hp + (*xilian).Defence
	}
}
func (xilian *Character) Skill(enemy *Character) { //技能
	(*enemy).Hp -= (*xilian).Attacks * 2
	if (*enemy).Hp <= 0 {
		(*enemy).IsAlive = false
		fmt.Println("成功击败")
	}
}
func (xilian *Character) Hurt(enemy *Character) { //受伤
	(*xilian).Hp -= (*enemy).Attacks
	if (*xilian).Hp <= 0 {
		(*xilian).IsAlive = false
	}
}
