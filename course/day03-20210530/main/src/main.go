package main

import "fmt"

var mainVar = "main Var"

func mainFunc() {
	fmt.Println("main func")
}
func main() {
	mainFunc()
	utilsFunc()
	fmt.Println(mainVar)
	fmt.Println(utilsVar)
}
