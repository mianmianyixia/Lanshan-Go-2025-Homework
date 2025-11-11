package main

import (
	Boss "123/homework3/boss"
	Character "123/homework3/character"
	Interface "123/homework3/interface"
	"fmt"
	"time"
)

func main() {
	var chose int
	var character *Interface.Character
	fmt.Println("Enter your choice:1.昔涟,2.敬请期待... ")
	fmt.Scan(&chose)
	if chose == 1 {
		character = &Character.PhLia093
	}
	boss := &Boss.Laigushi
	rounds(character, boss, character)
}

func rounds(character Interface.Action, boss *Interface.Character, characters *Interface.Character) {
	for round := 1; ; round++ {
		character.Die(characters)
		character.Die(boss)
		fmt.Printf("昔涟当前攻击力为:%d\n生命值为:%d\n防御力为%d\n技能点为%d\n", characters.Attack, characters.HP, characters.Defense, characters.Skill)
		time.Sleep(1 * time.Second)
		fmt.Printf("来古士当前攻击力为:%d\n生命值为:%d\n防御力为%d\n", boss.Attack, boss.HP, boss.Defense)
		time.Sleep(1 * time.Second)
		if !characters.IsAlive || !boss.IsAlive {
			if !characters.IsAlive {
				fmt.Println("昔涟暂时离场")
			} else {
				fmt.Println("来古士寄了")
			}
			return
		}
		fmt.Printf("round: %d\n", round)
		fmt.Println("请输入操作指令")
		fmt.Println("1.攻击 2.防御 3.技能")
		var control int
		fmt.Scan(&control)
		switch control {
		case 1:
			character.Attacks(boss)
		case 2:
			character.Defenses(boss)
		case 3:
			if !character.Skills(boss, &characters.Skill) {
				fmt.Println("致命空枪")
				time.Sleep(1 * time.Second)
			}
		default:
			fmt.Println("请输入正确的指令")
		}
		character.Hurt(boss)

	}
}
