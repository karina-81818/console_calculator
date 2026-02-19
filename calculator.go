package main

import (
	"fmt"
	"strconv"
)

func main(){

	leftOperand := inputNumber("Input left operand: ")
    operation := inputOperator("Input operation: ")
    rightOperand := inputNumber("Input right operand: ")

    res := calculate(leftOperand, rightOperand, operation)
	fmt.Println(res)
}

func inputNumber(str string) float64 {
	var input string

	for{
		fmt.Println(str)
		fmt.Scan(&input)
		num, err := strconv.ParseFloat(input, 0)
		if err == nil {
		return num
		} else {
		fmt.Println("Ошибка inputNumber: операнд должен быть числом. Введите верное число")
		}
	}
}

func inputOperator(str string) string {
    var operation string
    fmt.Println(str)
    fmt.Scan(&operation)

    if operation == "+" || operation == "-" || operation == "*" || operation == "/" {
        return operation
    }
    fmt.Println("Ошибка: введите верный оператор (+, -, *, /)")
    return inputOperator(str) 
}

func calculate(leftOperand, rightOperand float64, operation string) float64{
switch operation {
	case "+":
		return leftOperand + rightOperand
	case "-":
		return leftOperand - rightOperand
	case "*":
		return leftOperand * rightOperand
	case "/":
		if rightOperand == 0 {
        fmt.Println("Ошибка: деление на ноль!")
        return 0
    }
		return leftOperand / rightOperand
	default:
        fmt.Println("Ошибка: неизвестная операция")
        	return 0
	}
}
