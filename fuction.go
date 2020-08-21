package main

import (
	"strconv"
)

func main() {
	//fmt.Println(max(1,20))
	//fmt.Println(swap("sdd", "qiu"))
	//fmt.Println(contact("sdd", 1))

	//var c1 Circle
	//c1.radius = 10.00
	//fmt.Println("圆的面积 = ", c1.getArea())

	///* 声明函数变量 */
	//getSquareRoot := func(x float64) float64 {
	//	//return math.Sqrt(x)
	//	return math.Abs(x)
	//}
	///* 使用函数 */
	//fmt.Println(getSquareRoot(-9))

	/* nextNumber 为一个函数，函数 i 为 0 */
	//nextNumber := getSequence()
	//
	///* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//
	///* 创建新的函数 nextNumber1，并查看结果 */
	//nextNumber1 := getSequence()
	//fmt.Println(nextNumber1())
	//fmt.Println(nextNumber1())
}

func getSequence() func() int {
	i:=0
	return func() int {
		i+=1
		return i
	}
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func contact(a string, b int) string {
	return a + strconv.Itoa(b)
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

/* 定义结构体 */
type Circle struct {
	radius float64
}


