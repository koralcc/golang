package main

//fmt、math 是 Go 的标准库，它其实是去 GOROOT 下去加载该模块，当然 Go 的 import 还支持如下两种方式来加载自己写的模块：
import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外面(包内)定义变量不能用 :=，可以用var()定义多个

var aa = 11
var (
	ss = true
	bb = "abc"
)

func variableZeroValue() {
	//go语言语句结束不需要用分号;
	//go语言定义完变量即有初值
	//go语言定义后必须得用到，不然会报错

	var a int
	var s string
	//fmt.Println(a,s)//println不会打印出引号
	fmt.Printf("%d \n %q", a, s)
}

func triangle() {
	var a, b int = 3, 4
	//类型转换是强制的，没有自动转换
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func euler() {
	fmt.Printf("%3.f\n", cmplx.Exp(1i*math.Pi)+1)
}
func variableInitialValue() {
	//go语言可进行多个变量的初始化定义
	var a, b int = 1, 2
	var c string = "abc"
	fmt.Println(a, b, c)
}

func varibaleTypeDeduction() {
	//
	var a, b, c = 1, 2, "abc"
	fmt.Println(a, b, c)
}

func variableShorter() {
	//go语言块内变量可省略var和类型进行多个变量不同类型的初始化，但外面的变量不能用 :=进行定义，初始化声明和定义操作，不能对相同变量使用多次
	a, b, c := 1, 2, "abc"
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitialValue()
	varibaleTypeDeduction()
	variableShorter()
	triangle()
	euler()
	fmt.Println(aa, ss, bb)

}
