package main

import "fmt"

func factorial(count int) int {
	if count == 1 {
		return 1
	}
	return count * factorial(count-1)
}
func main() {
	fmt.Println("请输入数字：")
	var count int
	fmt.Scan(&count)
	sum := factorial(count)
	fmt.Println(sum)
}
