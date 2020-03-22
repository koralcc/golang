package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// for循环可以省略起始条件、递增语句、结束判断
// 什么都不加就是死循环 其他省略就类似于while，golang没有while关键字

//int转化为byte -- 省略初始条件
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		result += strconv.Itoa(n % 2)
	}

	return result

}

//一行一行读文件 -- 省略初始和递增条件
func printFile() {
	filename := "abc.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err) //终止程序并报系统错误
	}

	scanner := bufio.NewScanner(file)

	// = 此处完全等于while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环，省略初始条件、递增条件、结束条件,在后续的并发编程中重要的概念
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBin(5))
	printFile()
}
