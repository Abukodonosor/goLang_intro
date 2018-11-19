package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gde si decko !!")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi this is new route!!")

	var HTML string
	HTML = "<div style='width:200px;height:200px;background-color:orange;border:2px solid red;margin:100px auto;'>Welcome to world of GO!!</div>"
	fmt.Fprintf(w, "%s", HTML)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/home", newHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
