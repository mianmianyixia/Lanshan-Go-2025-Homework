package Interface

import "fmt"

// 获得角色信息
type Character struct {
	Name       string
	HP         float64
	MaxHP      float64
	Attack     float64
	BaseAttack float64
	Defense    float64
	CRT        float64
	IsAlive    bool
	Skill      int
	Energy     int
	BuffTime   int
}

// 用来执行行为
type Action interface {
	TakeDamage(float65 float64)
	StatusCheck()
}

func (people *Character) TakeDamage(damage float64) { //受击
	actualdamage := damage - (*people).Defense
	if actualdamage <= 0 {
		fmt.Println("未受到任何伤害")
		return
	} else {
		if people.HP < actualdamage {
			actualdamage = people.HP
		}
		(*people).HP -= actualdamage
		fmt.Printf("%s受到伤害为%.2f\n剩余生命为%.2f\n", (*people).Name, actualdamage, (*people).HP)
	}
}
func (people *Character) StatusCheck() {
	if (*people).HP == 0 {
		(*people).IsAlive = false
	}
	if !(*people).IsAlive {
		fmt.Printf("%s离场\n", (*people).Name)
	} else {
		fmt.Printf("%s当前攻击力为:%.2f\n生命值为:%.2f\n防御力为%.2f\n", (*people).Name, (*people).Attack, (*people).HP, (*people).Defense)
	}
	return
}
