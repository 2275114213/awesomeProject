package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "喜剧之王",
		Year:   2012,
		Price:  100,
		Actors: []string{"ldh", "hsy"},
	}

	// 将结构体转化成json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("=====%s\n", jsonStr)
	}

	// 把jsonStr 转成结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil{
		fmt.Println("=======",err)
	}
	fmt.Println("=",myMovie)

}
