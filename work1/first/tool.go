package main

import (
	"fmt"
	"strconv"
)

//Solve 调用的函数
func Solve(str string) (string, float32) {
	l, r := newMaths(str) //获取等式的左侧和右侧的每一个值
	move(&l, &r)
	return calculate(&l, &r)
}

//newMaths 根据输入的算式初始化math并返回math切片
func newMaths(str string) (l []*math, r []*math) {
	var repL []*math
	var repR []*math
	flag := false //是否已经遇到等号
	i := 0
	for ; i < len(str); i++ {
		symbol := "+"
		if str[i] == '=' {
			flag = true
		}

		//读取到符号后从此开始
	LABEL:
		//是常数则向后读取数字
		if isConst(str[i]) {
			j := 0
			var v string
			for ; i+j < len(str) && isConst(str[i+j]); j++ {
				v += string(str[i+j])
			}
			i += j - 1
			//向后取一位
			if i+1 < len(str) && isX(str[i+1]) {
				newMathWithFlag(symbol, v, string(str[i+1]), flag, &repL, &repR)
				i += 1
			} else {
				newMathWithFlag(symbol, v, "", flag, &repL, &repR)
			}
			//是未知数
		} else if isX(str[i]) {
			newMathWithFlag(symbol, "1", string(str[i]), flag, &repL, &repR)
			//是符号则向后判断
		} else if isSymbol(str[i]) {
			symbol = string(str[i])
			i++
			goto LABEL
		}
	}
	return repL, repR
}

//移项，将含未知数的项移到左边，常数项移到右边
func move(L *[]*math, R *[]*math) {
	i := 0
	for _, v := range *L {
		if v.x == "" {
			*R = append(*R, changeSymbol(*v))
			//删除这一项
			(*L)[i] = nil
		}
		i++
	}

	i = 0
	for _, v := range *R {
		if v.x != "" {
			*L = append(*L, changeSymbol(*v))
			//删除此项
			(*R)[i] = nil
		}
		i++
	}
}

//计算结果
func calculate(L *[]*math, R *[]*math) (string, float32) {
	v1, x := calculateX(*L)
	v2 := calculateV(*R)
	return x, float32(v2) / float32(v1)
}

//记录输入的数据
type math struct {
	symbol string
	value  int
	x      string
}

//translate 将数据转换为可以计算的int值，如果x存在只操作系数
func (m *math) translate() int {
	if m.symbol == "-" {
		return 0 - m.value
	} else {
		return m.value
	}
}

//输出结果便于修正
func (m *math) out() {
	if m == nil {
		return
	}
	if m.x == "" {
		fmt.Printf("(%s%v)", m.symbol, m.value)
	} else {
		fmt.Printf("(%s%v%s)", m.symbol, m.value, m.x)
	}
}

//根据标志f判断将数据加入哪一个切片中
func newMathWithFlag(s string, v string, x string, f bool, L *[]*math, R *[]*math) {
	if f {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		*R = append(*R, &math{symbol: s, value: i, x: x})
	} else {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		*L = append(*L, &math{symbol: s, value: i, x: x})
	}
}

//calculateX 计算x的参数并返回系数和
func calculateX(m []*math) (int, string) {
	rep := 0
	x := ""
	for _, v := range m {
		if v == nil {
			continue
		}
		if v.x == "" {
			panic("x的系数计算中必须包含x")
		}
		rep += v.translate()
		x = v.x
	}
	return rep, x
}

//calculateV 计算常数和
func calculateV(m []*math) int {
	rep := 0
	for _, v := range m {
		if v == nil {
			continue
		}
		if v.x != "" {
			panic("常数计算中不能包含x")
		}
		rep += v.translate()
	}
	return rep
}

//isSymbol 判断其是否为+或者—
func isSymbol(b byte) bool {
	if b == '-' || b == '+' {
		return true
	}
	return false
}

//isConst 判断其是否为数字
func isConst(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

//isX 判断其是否为未知数
func isX(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

func changeSymbol(m math) *math {
	if m.symbol == "+" {
		return &math{symbol: "-", value: m.value, x: m.x}
	} else {
		return &math{symbol: "+", value: m.value, x: m.x}
	}
}
