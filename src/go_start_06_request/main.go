package main

import (
	"fmt"
	"go_start_06_request/request"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Printf("Request scenario\n	")

	respons, err := http.Get("https://www.youtube.com/watch?v=eM2IXmfFWEM")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(respons.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)

	fmt.Println("======================== new way ====================")
	fmt.Println("=======================*********************====================")

	ReQuest := request.Request{}
	answer := ReQuest.Get("http://judo.rs/")

	fmt.Println("%s", answer)

}
