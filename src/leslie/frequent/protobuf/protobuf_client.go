package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	aaa "hello/leslie/frequent/protobuf/message"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/protobuf")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			user := &aaa.User{}
			proto.Unmarshal(body, user)
			fmt.Println(*user)
		}
	}
}
