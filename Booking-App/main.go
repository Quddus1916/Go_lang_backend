package main

import (
	"fmt"
	"net/http"
)

//get data from  a html form

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "error in parsing form %v", err)
		return
	}

	fmt.Fprintf(w, "post request successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name = %v \n", name)
	fmt.Fprintf(w, " address = %v", address)

}

// function handler
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

//simple server structure

func main() {
	filserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", filserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("server is running on port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server crashed")
	}

}
