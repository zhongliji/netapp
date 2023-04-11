package netapp

import (
	"fmt"
	"github.com/zhongliji/dell"
)

var Operator string
var Num1 float64
var Num2 float64
var C string

func Input() {

	fmt.Println("请输入操作符(+,-,*,/,%):")
	fmt.Scanln(&Operator)

	fmt.Println("请输入第一个数字:")
	fmt.Scanln(&Num1)

	fmt.Println("请输入第二个数字:")
	fmt.Scanln(&Num2)

	D := cal.Cal(Operator, Num1, Num2)
	fmt.Printf("%v %s %v = %v\n", Num1, Operator, Num2, D)

	fmt.Println("按c继续,按其他键退出")
	fmt.Scanln(&C)
}
