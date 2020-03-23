package main

import (
	"fmt"
)

//此处调数据时，方法里面会进行一个数组的拷贝
// [10]int和[20]int时不同的类型
func printArray(array [5]int) {
	array[3] = 10
	for i, v := range array {
		fmt.Println(i, v)
	}

}

//当然可以用指针的方式来改变数组值
func printArray2(array *[5]int) {
	array[3] = 10
	for i, v := range array {
		fmt.Println(i, v)
	}

}

//1.数组定义的三种方法
//2.数组是值类型，引用不会改变里面的值
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}          // :=定义数组时，必须初始化
	arr3 := [...]int{2, 4, 6, 8, 10} //不指定长度，系统自动算

	var grid [4][5]int //二维数组，表示，4个长度为5的int数组

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// for i, v := range arr1 {
	// 	fmt.Println(i, v)
	// }
	printArray(arr1)
	fmt.Println(arr1)
	printArray2(&arr1)
	fmt.Println(arr1)
}
