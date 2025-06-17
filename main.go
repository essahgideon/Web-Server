package main

import (
	"fmt"      //for standard text formatting
	"net/http" //for server tools ie. request and response
)

// The function that receives the client's request and submit's a response
func waiter(
	response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "my first Web Server. This is really interesting!")
}

// Main function
func main() {
	http.HandleFunc("/", waiter)
	//Confirmatory message to terminal
	fmt.Println("server is running on http://localhost:8080")

	//assigning the port to be used for the server
	http.ListenAndServe(":8080", nil)
}
