package Character

/*bug区：能量会超过上限
增益buff的持续时间
*/
import (
	Interface "123/homework3/interface"
	"fmt"
	"math/rand"
	"time"
)

// 初始化角色
// 用结构体在main函数中传入过多参数，考虑转成接口
var PhLia093 = Character{
	Interface.Character{
		Name:       "PhLia093",
		HP:         93,
		Attack:     13,
		BaseAttack: 13,
		Defense:    13,
		CRT:        0.5,
		IsAlive:    true,
		Skill:      3,
		Energy:     0,
		BuffTime:   0,
	}}

type Character struct {
	Interface.Character
}

func (xilan *Character) SkillDescription() {
	fmt.Print("使用普攻造成100%攻击力的伤害\n使用战技1，增强50%攻击力，持续三回合\n使用战技2，为自身回复等同攻击力的生命值\n当能量达到100时，昔涟可激活终结技，造成固定50的伤害\n")
}

// xilianAct是填入boss的接口
func (xilian *Character) NormalAttack(xilianAct Interface.Action) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < (*xilian).CRT {
		fmt.Println("出暴击了")
		damage := (*xilian).Attack * 1.5
		xilianAct.TakeDamage(damage)
	} else {
		fmt.Println("未出暴击")
		damage := (*xilian).Attack
		xilianAct.TakeDamage(damage)
	}
	(*xilian).Skill++
	(*xilian).Energy += 20
	if (*xilian).Skill > 5 {
		(*xilian).Skill = 5
	}
}
func (xilian *Character) Skill_1() bool {
	if (*xilian).Skill < 1 {
		fmt.Println("战技点不足")
		return false
	}
	(*xilian).Skill--
	(*xilian).Energy += 30
	fmt.Println("盛放吧,来世的乐土")
	if (*xilian).BuffTime == 0 {
		(*xilian).BuffTime = 3
		(*xilian).Attack = (*xilian).BaseAttack * 1.5
		return true
	} else {
		(*xilian).BuffTime = 3
		return true
	}
}

func (xilian *Character) Skill_2() bool {
	if (*xilian).Skill < 1 {
		fmt.Println("战技点不足")
		return false
	}
	(*xilian).Skill--
	(*xilian).Energy += 30
	fmt.Println("此诗，献予一切生命")
	(*xilian).HP += (*xilian).Attack
	return true
}
func (xilian *Character) Finisher(xilianAct Interface.Action) bool {
	if (*xilian).Energy == 100 {
		(*xilian).Energy = 0
		fmt.Println("花与箭之舞")
		damage := 50.0
		xilianAct.TakeDamage(damage)
		return true
	}
	fmt.Println("能量不足")
	return false
}
func (xilian *Character) Check(xilianAct Interface.Action) { //此处填角色
	xilianAct.StatusCheck()
	if (*xilian).IsAlive {
		fmt.Printf("战技点%d\n能量%d\n技能一buff持续时间%d\n", (*xilian).Skill, (*xilian).Energy, (*xilian).BuffTime)
	}
}
