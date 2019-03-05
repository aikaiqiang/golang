package main

import (
	"fmt"
	"strconv"
)

//行数
const LINES int = 10

func main() {
	fmt.Println("Hello, Go!")
	fmt.Printf("字符串转Unicode:\n")
	strToUnicode()

	fmt.Printf("杨辉三角:\n")
	ShowYangHuiTriangle()

	fmt.Printf("九九乘法表：\n")
	multiplicationTable()
}

func strToUnicode() {
	var str = "少年强则中国强"
	var b []byte = []byte(str)
	var chars []rune = []rune(str) //将str转换为rune数组（[]rune）
	fmt.Printf("b =%v, len(str)=%d\n", b, len(str))
	fmt.Printf("%c\n", 97)
	fmt.Printf("chars =%v, chars count:%d\n", chars, len(chars))
}

// 杨辉三角
func ShowYangHuiTriangle() {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print("  ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, "  ")
		}
		fmt.Println("")
	}
}

func add(a, b int) int {
	return a + b
}

// 九九乘法表
func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			var ret string
			if i*j < 10 && j != 1 {
				ret = " " + strconv.Itoa(i*j)
			} else {
				ret = strconv.Itoa(i * j)
			}

			fmt.Print(j, " * ", i, " = ", ret, "   ")
		}
		fmt.Print("\n")
	}
}
