package main

import (
	"fmt"
	"io/ioutil"
)

//switch操作，switch可以不用写任何表达式，在case里面进行判断，类似于if
func grade(score int) string {
	g := ""

	switch {
	case score < 60:
		g = "F" //golang 自动break,除非使用fallthrough
	case score < 70:
		g = "E"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		// panic(fmt.Sprintf("Wrong score : %d", score)) //`Sprintf` 格式化并返回一个字 符串而不带任何输出,printf会在屏幕打印出
		panic("error") //panic中断程序，并且打印出错误
	}

	return g
}

func main() {
	const (
		filename = "abc.txt"
	)

	//if里面可以赋值变量，在if范围里面有效
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(59),
		grade(83),
		grade(98),
		grade(101),
	)
}

// fmt.Printf(a,s)   Printf 可以打印字符串
// fmt.Printf("%d %s\n",a,s)  %s 不打印空字符串
// fmt.Printf("%d %q\n",a,s)  %q 打印空字符串
// fmt.Println(a,s)   Println 只能跟变量的名字
