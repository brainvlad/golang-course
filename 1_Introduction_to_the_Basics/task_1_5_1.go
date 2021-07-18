package main

import "fmt"

func main(){
	var a int
	fmt.Scan(&a)
	// Число умножается на 2;
	// Затем к числу прибавляется 100.
	a = a * 2 + 100
	fmt.Println(a)
}