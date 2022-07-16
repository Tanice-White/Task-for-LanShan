package main

import "fmt"

//Node 创建地图的时候储存F和C的位置
type Node struct {
	x int
	y int
}

type Solve struct{}

//NewMap 创建地图并获取F和C的位置
func (s *Solve) NewMap() ([10][10]string, Obj, Obj) {
	var (
		theMap [10][10]string //最终生成的地图
		strGet string         //获取到的输入
		_F     Obj            //F的位置
		_C     Obj            //C的位置
	)

	//初始化地图
	for i := 0; i < 10; i++ {
		_, _ = fmt.Scanln(&strGet)
		for k, v := range strGet {
			if k == 10 {
				panic("一行最多十个字符哦")
			}
			//初始化F和C的位置
			if v == 'F' {
				_F = NewObj(k, i)
			}
			if v == 'C' {
				_C = NewObj(k, i)
			}
			theMap[i][k] = string(v)
		}
	}
	return theMap, _F, _C
}

//out 可视化地图
func (s *Solve) out(m *[10][10]string) {
	for _, v := range m {
		for _, _v := range v {
			fmt.Printf("%s ", _v)
		}
		fmt.Println()
	}
}

//Obj 表示可移动的事物
type Obj struct {
	Node              //表示此物的位置
	way       int     //表示取direction特殊位置的值
	direction [4]byte //表示移动的位置
	Minute    int     //已经花费的时间
}

//NewObj 初始化
func NewObj(x int, y int) Obj {
	d := [4]byte{'w', 'd', 's', 'a'} //按照顺时针顺序
	n := Node{x: x, y: y}
	return Obj{
		Node:      n,
		way:       0,
		direction: d,
		Minute:    0,
	}
}

//changeWay 转向
func (o *Obj) changeWay() {
	o.way++
	if o.way == 4 {
		o.way = 0
	}
}

//move 表示物体的移动
//返回自己便于自调用
func (o *Obj) move(m *[10][10]string) {
	if o.direction[o.way] == 'w' {
		//若已经到顶部或者遇到障碍物则花费一分钟转向
		if o.y == 0 || (*m)[o.y-1][o.x] == "*" {
			o.changeWay()
		} else {
			o.y--
		}
	} else if o.direction[o.way] == 's' {
		if o.y == 9 || (*m)[o.y+1][o.x] == "*" {
			o.changeWay()
		} else {
			o.y++
		}
	} else if o.direction[o.way] == 'a' {
		if o.x == 0 || (*m)[o.y][o.x-1] == "*" {
			o.changeWay()
		} else {
			o.x--
		}
	} else if o.direction[o.way] == 'd' {
		if o.x == 9 || (*m)[o.y][o.x+1] == "*" {
			o.changeWay()
		} else {
			o.x++
		}
	} else {
		panic("运动位置错误")
	}
	o.Minute++
}

func (o *Obj) out() {
	fmt.Printf("message: \n[%v][%v] 时间：%v\n", o.x, o.y, o.Minute)
}

//IsOver 判断两人的位置是否重合
func IsOver(f *Obj, c *Obj) bool {
	if f.x == c.x && f.y == c.y {
		return true
	}
	return false
}
