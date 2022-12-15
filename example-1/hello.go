// Para crear un módulo nuevo (go.mod) se ejecuta "go mod init path-o-repo-name/mod-name"
// Paquete principal
package main

//Importamos paquetes externos: para instalarlos ejecutar "go mod tidy"
import (
	"fmt"
	"math/rand"
	"time"

	"rsc.io/quote/v4"
)

//Función principal del programa
func main(){
	rand.Seed(time.Now().Unix())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(rand.Intn(10))
	fmt.Println("¡Que tengan buen código!")
}