package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int
	_ = age
	age = 8900
	amount := 12.09
	name := "数据执行hhhhhhhhhhhhh"
	fmt.Println(amount)
	fmt.Println(name + "-->Hello Word!-->" + strconv.Itoa(age))
}
