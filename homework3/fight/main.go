package main

import (
	Boss "123/homework3/boss"
	Character "123/homework3/character"
	"fmt"
	"time"
)

func main() {
	var chose int
	var character *Character.Character
	fmt.Println("Enter your choice:1.昔涟,2.敬请期待... ")
	fmt.Scan(&chose)
	if chose == 1 {
		character = &Character.PhLia093
	}
	boss := &Boss.Laigushi
	rounds(character, boss)
	if !(*character).IsAlive {
		fmt.Println("挑战失败")
	} else {
		fmt.Println("挑战成功")
	}
}

func rounds(character *Character.Character, boss *Boss.Boss) {
	character.Check(character)
	boss.Check(boss)
	for round := 1; ; round++ {
		if !(*character).IsAlive || !(*boss).IsAlive {
			return
		}
		fmt.Println("轮到我方行动")
		characteract(character, boss, round)
		if !(*character).IsAlive || !(*boss).IsAlive {
			return
		}
		time.Sleep(3 * time.Second)
		fmt.Println("轮到敌方行动")
		bossact(boss, character)
		time.Sleep(3 * time.Second)
	}
}

func characteract(character *Character.Character, boss *Boss.Boss, round int) {
	for {
		if round > 1 {
			if (*character).BuffTime > 0 {
				(*character).BuffTime--
			} else {
				(*character).Attack = (*character).BaseAttack
			}
		}
		fmt.Printf("请输入操作\n1.普通攻击\t2.战技1\t3.战技2\t4.终结技\t5.查看角色技能描述\n")
		var control int
		fmt.Scan(&control)
		switch control {
		case 1:
			character.NormalAttack(boss)
			if (*character).Energy > 100 {
				(*character).Energy = 100
			}
			boss.Check(boss)
			return
		case 2:
			if character.Skill_1() {
				if (*character).Energy > 100 {
					(*character).Energy = 100
				}
				character.Check(character)
				return
			}
		case 3:
			if character.Skill_2() {
				if (*character).Energy > 100 {
					(*character).Energy = 100
				}
				character.Check(character)
				return
			}
		case 4:
			if character.Finisher(boss) {
				boss.Check(boss)
				return
			}
		case 5:
			character.SkillDescription()
		default:
			fmt.Println("请输入正确的操作")
		}
	}
}

func bossact(boss *Boss.Boss, character *Character.Character) {
	if (*boss).HP >= (*boss).MaxHP/2.0 {
		boss.NormalAttack(character)
		character.Check(character)
	} else if (*boss).HP >= (*boss).MaxHP/4.0 {
		if (*boss).BuffTime != 1 {
			boss.SecStage()
			boss.Check(boss)
		}
		boss.Skill_1(character)
		character.Check(character)
	} else {
		boss.Finisher(character)
		character.Check(character)
	}
}
