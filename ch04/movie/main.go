package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // 这是结构体成员的 tag ，可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列
	Color  bool `json:"color,omitempty"` // omitempty 表示当Go语言结构体成员为空或零值时不生成该JSON对象
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {

	data, err := json.MarshalIndent(movies, "", "   ") // struct --> json 的过程叫 marshaling
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

}
