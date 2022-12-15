package main

import "fmt"

func add(num1, num2 int) (int){
	return num1 + num2
}

func substract(num1, num2 int) (int){
	return num1 - num2
}

//Múltiples resultados/retornos
func multiply(num1, num2 int) (int, int, int){
	return num1 * num2, num1, num2
}

//"naked result" nombramos el valor retornado como "divisionResult" y el return queda vacío 
func divide(num1, num2 int) (divisionResult int){
	divisionResult = num1 / num2
	return
}

func main(){
	fmt.Println(add(1,2))
	fmt.Println(substract(1,2))
	var multiplyResult, firstNumber, secondNumber = multiply(1,2)
	fmt.Println("Multiply", multiplyResult, firstNumber, secondNumber)
	fmt.Println(divide(1,2))
}