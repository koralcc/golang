package main

import (
	"fmt"
)

func printslice(s []int) {
	fmt.Printf("%v,len =%d , cap=%d \n", s, len(s), cap(s))
}
func main() {
	//
	var s []int
	for i := 0; i < 100; i++ {
		printslice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printslice(s1)

	fmt.Println("Makeing slice")
	s2 := make([]int, 5)
	s3 := make([]int, 8, 32)
	printslice(s2)
	printslice(s3)
	fmt.Println("Copying slice")
	copy(s2, s1)
	printslice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) //可变参数可以通过slice...实现
	printslice(s2)

	fmt.Println("Poping from front")
	s2 = s2[1:]
	printslice(s2)

	fmt.Println("Poping from back")
	s2 = s2[:len(s2)-1]
	printslice(s2)
}
