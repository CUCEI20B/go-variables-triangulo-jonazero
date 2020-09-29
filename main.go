package main

import "fmt"

func main(){
	var base float64
	var altura float64
	var area float64
	fmt.Scanln(&base)
	fmt.Scanln(&altura)
	area = (base * altura)/2
	fmt.Println(area)
}