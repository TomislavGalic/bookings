package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main app function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.Abt)

	fmt.Println(fmt.Sprintf("Starting app on port: %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
