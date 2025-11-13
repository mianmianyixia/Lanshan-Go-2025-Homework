package Boss

import (
	Interface "123/homework3/interface"
	"fmt"
	"math/rand"
	"time"
)

// boss阶段性，放技能的随机性
type Boss struct {
	Interface.Character
}

var Laigushi = Boss{
	Interface.Character{
		Name:     "来古士",
		HP:       78,
		MaxHP:    78,
		Attack:   7,
		Defense:  8,
		IsAlive:  true,
		BuffTime: 0,
	}}

// BossAct攻击对象为角色
func (boss *Boss) NormalAttack(BossAct Interface.Action) {
	damage := (*boss).Attack
	BossAct.TakeDamage(damage)
}
func (boss *Boss) SecStage() {
	fmt.Println("论证形态")
	(*boss).BuffTime = 1
	(*boss).Attack *= 2

}
func (boss *Boss) Skill_1(BossAct Interface.Action) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.3 {
		fmt.Println("毁灭论证法")
		damage := (*boss).Attack * 2
		BossAct.TakeDamage(damage)
	} else {
		boss.NormalAttack(BossAct)
	}
}
func (boss *Boss) Finisher(BossAct Interface.Action) {
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < 0.1 {
		fmt.Println("求解:第一因之果")
		damage := (*boss).Attack * 6
		BossAct.TakeDamage(damage)
	} else {
		boss.Skill_1(BossAct)
	}
}
func (boss *Boss) Check(BossAct Interface.Action) {
	BossAct.StatusCheck()
}
