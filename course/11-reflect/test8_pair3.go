package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Write interface {
	WriteBook()
}

type Book2 struct {
}

func (t *Book2) ReadBook() {
	fmt.Println("Read a book")
}

func (t *Book2) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	// b pair<type:*Book2,value:{}>
	b := &Book2{}
	// r pair<type: value>
	var r Reader
	// r pair<type:*Book2,value:{}>
	r = b
	r.ReadBook()

	var w Write
	w = r.(Write)  // 因为W 和 r 具体的type 是一致的
	w.WriteBook()
}
