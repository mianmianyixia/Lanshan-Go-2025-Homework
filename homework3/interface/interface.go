package Interface

import "fmt"

// 获得角色信息
type Character struct {
	Name    string
	HP      int
	Attack  int
	Defense int
	IsAlive bool
	Skill   int
}

// 用来执行行为
type Action interface {
	Attacks(*Character)
	Defenses(*Character)
	Skills(*Character, *int) bool
	Die(*Character)
	Hurt(*Character)
}

func (character *Character) Attacks(enemy *Character) {
	(*enemy).HP = (*enemy).HP - (*character).Attack + (*enemy).Defense
}
func (character *Character) Defenses(enemy *Character) {
	if (*character).Defense < (*enemy).Attack {
		(*character).HP = (*character).Defense + (*character).HP
	} else {
		fmt.Println("敌人未造成任何伤害")
		(*character).HP = (*enemy).Attack + (*character).HP
		(*enemy).HP -= (*character).Attack / 2 //防御成功会反击
	}
}
func (character *Character) Skills(enemy *Character, num *int) bool {
	if *num == 0 {
		fmt.Println("技能点不够了")
		return false
	}
	*num--
	(*enemy).HP -= (*character).Attack*2 - (*enemy).Defense
	return true

}
func (character *Character) Die(characters *Character) {
	if (*characters).HP <= 0 {
		(*characters).IsAlive = false
		(*characters).HP = 0
	}
	if !(*characters).IsAlive {
		fmt.Println("战斗结束")
		return
	}
}
func (character *Character) Hurt(enemy *Character) {
	(*character).HP -= (*enemy).Attack
}
