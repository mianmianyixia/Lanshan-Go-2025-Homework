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
	var mark string
	fmt.Printf("欢迎使用Go语言计算器\n")

	fmt.Printf("请输入一个整数:")
	fmt.Scan(&num1)
	for i := 1; ; i++ {
		fmt.Printf("请输入操作符:")
		fmt.Scan(&mark)
		if mark == "=" {
			break
		}
		fmt.Printf("请输入一个整数:")
		fmt.Scan(&num2)
		if i == 1 { //为第一次的sum赋值
			sum = num1
		}
		sum = jisuan(control(mark), sum, num2)
	}
	fmt.Println(sum)

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

//可用结构体
