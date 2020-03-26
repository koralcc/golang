package main

import (
	"fmt"
)

//[]int 代表一个slice参数,slice本身没有数据，是对底层数组的一个view
func updateSlices(sl []int) {
	sl[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[0:2]
	s2 := arr[:4]
	s3 := arr[3:]
	s4 := arr[:]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)
	updateSlices(s2[:])
	fmt.Println(s2)
	fmt.Println(s1)
	fmt.Println(arr)

	//reslice，对切片自身的一个slice
	//ptr,slice的第一个
	//len,slice的长度，下标不能超过这个值
	//cap,容量，向后扩展不能超过这个值
	fmt.Println("Reslice")
	fmt.Println(s4)
	s4 = s4[0:3]
	fmt.Println(s4)
	s4 = s4[2:4]
	fmt.Println(s4)
	s4 = s4[2:6]
	fmt.Println(s4)
	fmt.Printf("s1 = %v len=%d cap = %d", s4, len(s4), cap(s4))
	//添加元素时，如果超越了cap，系统会重新分配更大的底层数组
	//由于值传递的关系，必须接收append的返回值
	s5 := append(s4, 11)
	s6 := append(s5, 15)
	s7 := append(s6, 16)
	fmt.Println("Append")
	fmt.Println(s4)
	fmt.Println(arr)
	fmt.Printf("s5, s6, s7 =", s5, s6, s7)
}
