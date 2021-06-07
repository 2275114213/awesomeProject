package main

import "fmt"

func test20() (err error) {
	defer func() {
		fmt.Println("defer")
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("%s", panicError)
		}
	}()
	fmt.Println("before")
	// raise Error
	panic("自定义panic")
	fmt.Println("after")
	return
}

func main() {
	fmt.Println("before main")
	err := test20()
	fmt.Println(err)

}
