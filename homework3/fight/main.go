// 小登好多东西都用不明白，只有先做个残缺半成品，先把作业交了，这两天再想想，给它完善一下
package main

import (
	"123/homework3/PhLia093"
	"fmt"
)

func main() {
	var action string
	var xilian = &PhLia093.Character{
		Attacks: 10,
		Defence: 13,
		Hp:      93,
		IsAlive: true,
	}
	laigushi := &PhLia093.Character{
		Attacks: 7,
		Hp:      78,
		IsAlive: true,
	}
	characteraction(xilian, laigushi, xilian, action)
}

// 传入角色boss数据和行为及指令
func characteraction(characters PhLia093.Action, boss *PhLia093.Character, character *PhLia093.Character, control string) {
	fmt.Println("战斗开始")

	for i := 1; ; i++ {
		if !(*boss).IsAlive || !(*character).IsAlive { //
			fmt.Println("战斗结束")
			return
		}
		fmt.Printf("昔涟当前攻击力为:%3.f\n生命值为:%3.f\n防御力为%3.f\n", character.Attacks, character.Hp, character.Defence)
		fmt.Printf("来古士当前攻击力为:%3.f\n生命值为:%3.f\n", boss.Attacks, boss.Hp)
		fmt.Printf("第%d回合", i)
		fmt.Printf("请操作你的角色\n1.攻击\n2.技能\n3.防御\n")
		fmt.Scan(&control)
		for {
			if control == "攻击" {
				characters.Attack(boss)
				break
			} else if control == "技能" {
				characters.Skill(boss)
				break
			} else if control == "防御" {
				characters.Defend(boss)
				break
			} else {
				fmt.Println("请输入正确的操作")
			}
		}
		characters.Hurt(boss)
		if !(*boss).IsAlive {
			fmt.Println("来古士挂了")
		}
		if !(*character).IsAlive {
			fmt.Println("昔涟离场了")
		}
	}
}
