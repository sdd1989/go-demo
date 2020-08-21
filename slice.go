package main

import "fmt"

func main() {
	//var numbers = make([]int, 3, 5)
	//numbers = append(numbers, 1,2,3)
	//printSlice(numbers)
	//numbers2 := make([]int, len(numbers), cap(numbers) * 2)
	//printSlice(numbers2)
	//copy(numbers2, numbers)
	//printSlice(numbers)
	//var arr = []int{1,2,3,4,5}
	//printSlice(arr)
	//var array = [...]int{111,222,333,444,555}
	//numbers := array[1:3]
	//printSlice(numbers)
	//数组
	var array = [...]int{1,2,3}
	var array2 = [3]int{1,2,3}
	fmt.Println("数组:", array,array2)
	//切片
	var slice = []int{1,2,3}
	var slice2 = make([]int, 3, 10)
	var slice3 = array[:]
	fmt.Println("切片:", slice,slice2,slice3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x), cap(x), x)
}