package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	testJson1()
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`          // 成员标签. 将Struct中的Year, 转成json中的released
	Color  bool `json:"color,omitempty"`   // 成员标签, 将Struct中的Color, 转成json中的color, 忽略零值或空.
	Actors []string
}

func testJson1() {
	fmt.Println("======test1=====")
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON Marshaling failed.")
	}
	fmt.Printf("%s\n", data)
	data, _ = json.MarshalIndent(movies, "", "   ")
	fmt.Printf("%s\n", data)

	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalln("error")
	}
	fmt.Printf("titles: %s\n", titles)
}
