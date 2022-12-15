package main

import "fmt"

//declarar variables fuera de funciones con var

var name, lastname string

//inicializar variables con tipado implícito
var phoneNumber = "12345667"


func returnData()(string, bool, float64){
	//conversión de tipos: T(v) convierte v a tipo T
	const number int = 3
	var converted float64 = float64(number)
	works := true
	//declaración "corta"
	address := "dirección"
	return address, works, converted
}

func main(){
	fmt.Println(returnData())
}