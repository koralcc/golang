package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数返回多个值时可以起名字，但当返回的参数太多时不建议取名字，对调用者是没有区别的
//多个值返回，比较常见的场景是返回error
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsupported operation :" + op)
		return 0, fmt.Errorf("unsupported operation :" + op)
	}
}

func div(a int, b int) (q, r int) {
	return a / b, a % b
}

//golang函数式编程，可以将方法当参数传入，传入方法名字，入参和出参
func apply(op func(a, b int) int, a, b int) int {
	//利用反射获取获取方法的信息
	p := reflect.ValueOf(op).Pointer()
	//runtime获取方法名
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//go支持可变参数列表,当数组用即可
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

// 1.go语言指针不能运算
//
// 2.参数传递——值传递？引用传递？
// go语言只有值传递一种方式，——拷贝参数
//
// 3.指针方式实现引用传递
//go 指针，go语言只有值传递
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	if result, err := eval(3, 4, "X"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
