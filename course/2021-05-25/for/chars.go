package main

import "fmt"

func main() {
	letters := "dsdfdfdfdf"
	//for i := 1; i < len(letters); i++ {
	//	fmt.Printf("%c\n", letters[i])
	//}
	for i := 0; i < len(letters); i++ {
		fmt.Printf("%c\n", letters[i])
	}
	msg := "我爱祖国"
	for _, v := range msg {
		fmt.Printf("%q\n", v)
	}

}
