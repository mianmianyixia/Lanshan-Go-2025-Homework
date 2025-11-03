package main

import "fmt"

type operate func(int, int) int

func control(mark string) operate {

	switch mark {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	default:
		return nil
	}
}
func jisuan(control operate, num1, num2 int) int {
	return control(num1, num2)
}

func main() {
	var num1, num2, sum int
	var action, mark string
	fmt.Printf("欢迎使用Go语言计算器\n请输入两个整数和一个运算符，进行四则运算\n输入exit退出\n")
	for {
		fmt.Printf("请输入第一个整数:")
		fmt.Scan(&num1)
		fmt.Printf("请输入操作符:")
		fmt.Scan(&mark)
		fmt.Printf("请输入第二个整数:")
		fmt.Scan(&num2)
		sum = jisuan(control(mark), num1, num2)
		fmt.Println(sum)
		fmt.Printf("是否继续?(exit退出)")

		for {
			fmt.Scan(&action)
			if action == "exit" {
				fmt.Println("感谢使用!再见!")
				return
			} else if action != "y" {
				fmt.Println("请输入正确的控制")
			} else {
				break
			}
		}
	}
}
func add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}
func sub(num1, num2 int) int {
	sum := num1 - num2
	return sum
}
func mul(num1, num2 int) int {
	sum := num1 * num2
	return sum
}
func div(num1, num2 int) int {
	sum := num1 / num2
	return sum
}
