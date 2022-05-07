package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func Index(w http.ResponseWriter, r *http.Request){
	fmt.Println(w, "Blog:aaaa")
}
