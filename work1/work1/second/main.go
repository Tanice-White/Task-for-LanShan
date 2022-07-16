package main

import "fmt"

func main() {
	fmt.Println("请输入地图：")
	Solution := Solve{}
	m, F, C := Solution.NewMap()
	for !IsOver(&F, &C) {
		F.move(&m)
		C.move(&m)
		if F.Minute > 10000 {
			fmt.Println("0")
			return
		}
	}
	fmt.Println(F.Minute)
}
