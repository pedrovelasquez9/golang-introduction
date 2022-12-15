package main

import "fmt"

//for loop
func accum() (int){
	helper:=0
	for i:=0; i<10; i++{
		helper += i
	}
	return helper
}

//while loop

func sumatory()(int){
	sum := 0
	for sum < 10{
		sum++
	}
	return sum
}

//if conditional

func amIAlive(actualCodingHours int) (string){
	maxCodingHours := 12
	if actualCodingHours < maxCodingHours {
		return "Staying alive"
	}
	return "Staying aliiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiive"
}

func main(){
	fmt.Println(accum())
	fmt.Println(sumatory())
	fmt.Println(amIAlive(12))
}