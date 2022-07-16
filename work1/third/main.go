package main

import "fmt"

func main() {
	var (
		n    int    //偏移量
		pStr string //原密码
	)
	fmt.Println("请输入n：")
	_, _ = fmt.Scanln(&n)
	fmt.Println("请输入原字符串：")
	_, _ = fmt.Scanln(&pStr)
	fmt.Println("密码是", deviation(pStr, n))
}

func deviation(str string, n int) string {
	var rep string
	for _, v := range str {
		if n > 26 {
			n -= 26
		}
		v += int32(n)
		rep += string(v)
	}
	return rep
}
