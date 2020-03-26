package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println(m)

	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)
	fmt.Println("Tranversing map")

	//遍历 range，不保证顺序，可遍历k，也可遍历k ，v
	for k, v := range m {
		fmt.Println(k, v)
	}

	//key不存在时，也不报错，可用ok来接
	courseName := m["course"]
	fmt.Println(courseName)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除值，用delete方法
	delete(m, "name")
	name, ok := m["name"]
	fmt.Println(name, ok)
	//map的key 初 slice map function外的都可以，struct类型不包含以上字段时，也可做为key
}
