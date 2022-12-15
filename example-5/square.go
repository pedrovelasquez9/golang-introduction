// simular la función de raíz cuadrada dentro de un loop para aproximar el valor recibido a través de la operación z -= (z*z - x) / (2*z)
package main

import (
	"fmt"
	"math"
)

var z float64 = 1.0
var i int = 0
var y float64 = 0

func Sqrt(x float64) (float64, int) {
	for getPositiveValue(z-y) > 2.3e-16 {
		y = z
		z -= (z*z - x) / (2*z)
		i++
	}
	return z,i
}

func getPositiveValue(num float64) (float64){
	if(num < 0){
		return num * -1
	}
	return num
}

func main() {
	fmt.Println(math.Sqrt(32357))
	var result, iterations = Sqrt(32357);
	fmt.Println(result, iterations)
}
