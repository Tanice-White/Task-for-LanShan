//数字拆分1<=n<=8
package main

import "fmt"

var a [100]int = [100]int{1}
var n int

func main() {
	_, _ = fmt.Scanln(&n)
	search(n, 1)
}

func search(s int, t int) {
	var i int
	for i = a[t-1]; i <= s; i++ {
		if i < n {
			a[t] = i
			s -= i
			if s == 0 {
				myPrint(t) //当s=0时，拆分结束输出结果
			} else {
				search(s, t+1)
			}
			s += i //回溯：加上拆分的数，以便产生所有可能的拆分
		}
	}
}

func myPrint(t int) {
	for i := 1; i <= t-1; i++ {
		fmt.Print(a[i], "+")
	}
	fmt.Println(a[t])
}
