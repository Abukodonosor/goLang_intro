package main

// https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gde si decko !!")
	//r.URL.Query().Get("keyword")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi this is new route!!")

	var HTML string
	HTML = "<div style='width:200px;height:200px;background-color:orange;border:2px solid red;margin:100px auto;'>Welcome to world of GO!!</div>"
	fmt.Fprintf(w, "%s", HTML)
}

func manuHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Usao u menuHandler!!!")
	data, err := ioutil.ReadFile("./index.html")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(data))
}

//midleware
// info https://www.alexedwards.net/blog/making-and-using-middleware
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func main() {
	// serve static content
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// declare routes of webApp
	go http.HandleFunc("/", handler)
	go http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/menu", manuHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
