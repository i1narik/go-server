package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(writer, "Parse successfull\n")
	name := request.FormValue("name")
	age := request.FormValue("age")
	fmt.Fprintf(writer, "Name: %s\nAge: %s", name, age)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "Not allowed method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(writer, "hello")
}
