package main

import "fmt"

func average(sum int, count int) float64 {
	var average float64
	average = float64(sum) / float64(count)
	return average
}

func main() {
	var num, sum, count int
	for {
		fmt.Println("请输入一个整数：")
		fmt.Scanln(&num)
		sum += num
		if num == 0 {
			break
		}
		count++
	}
	average := average(sum, count)
	fmt.Printf("平均成绩为%.2f,", average)
	if average >= 60 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格")
	}
}
