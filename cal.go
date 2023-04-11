package cal

import "fmt"

func Cal(Operator string, Num1 float64, Num2 float64) float64 {

	var Result float64
	switch Operator {
	case "+":
		Result = Num1 + Num2
	case "-":
		Result = Num1 - Num2
	case "*":
		Result = Num1 * Num2
	case "/":
		Result = Num1 / Num2
	case "%":
		Result = float64(int(Num1) % int(Num2))
	default:
		fmt.Println("你输入的操作符不正确,支持的操作符有:+,-,*,/,%")

	}

	return Result

}
