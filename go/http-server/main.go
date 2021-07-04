package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside homeHandler...")
	/* Write text to web page */
	fmt.Fprintf(w, "Hi, This is from Home Handler")
}

func page1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside page1Handler...")
	fmt.Fprintf(w, "Hello, This is from Page 1 Handler")
}

func page2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside page2Handler...")
	fmt.Fprintf(w, "Hello, This is from Page 2 Handler")
}

func main() {
	/* Add handler */
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/page1", page1Handler)
	http.HandleFunc("/page2", page2Handler)

	/* Listen on a port */
	fmt.Println("Listening at 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ERROR: ListenAndServe:", err)
	}
}
