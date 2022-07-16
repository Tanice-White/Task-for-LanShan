package main

import "fmt"

func main() {
	var formula string
	fmt.Println("请输入算式：")
	_, _ = fmt.Scanln(&formula)
	x, rep := Solve(formula)
	fmt.Printf("%s = %0.3f", x, rep)
}
