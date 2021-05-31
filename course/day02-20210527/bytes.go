package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Compare([]byte("dsdsd"), []byte("sdsd")))
}
